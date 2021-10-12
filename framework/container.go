package framework

import (
	"errors"
	"fmt"
	"sync"
)

// Container 是一个服务容器，提供绑定服务和获取服务的功能
type Container interface {
	// Bind 绑定一个服务提供者，如果关键字凭证已经存在，会进行替换操作，不返回error
	Bind(provider ServiceProvider)error
	// IsBind 关键字凭证是否已经绑定服务提供者
	IsBind(key string) bool
	// Make 根据关键字凭证获取一个服务
	Make(key string) (interface{},error)
	// MustMake 根据关键字凭证获取一个服务,如果这个关键字凭证未绑定服务提供者，那么会panic
	// 所以在使用该接口的时候，请保证服务容器已经为这个关键字凭证绑定服务提供者
	MustMake(key string) interface{}
	// MakeNew 根据关键字凭证获取一个服务，只是这个服务并不是单例模式的
	// 它是根据服务提供者注册的启动函数和传递的params参数实例化出来的
	// 这个函数在需要为不同参数启动不同实例的时候非常有用
	MakeNew(key string,params ...interface{}) (interface{},error)
}

type ZettContainer struct {
	Container
	// providers 存储注册的服务提供者，key为字符串凭证
	providers map[string]ServiceProvider
	// instances 存储具体的实例，key为字符串凭证
	instances map[string]interface{}
	// lock 用于锁住对容器的变更操作
	lock sync.RWMutex
}
func NewZettContainer() Container{
	return &ZettContainer{
		providers:map[string]ServiceProvider{},
		instances: map[string]interface{}{},
		lock:sync.RWMutex{},
	}
}

func (zett *ZettContainer) Bind(provider ServiceProvider) error{
	zett.lock.Lock()
	defer zett.lock.Unlock()
	key := provider.Name()
	zett.providers[key] = provider
	if !provider.IsDefer(){
		params := provider.Params(zett)
		if err := provider.Boot(zett);err != nil{
			return err
		}
		method := provider.Register(zett)
		instance,err := method(params...)
		if err != nil{
			return err
		}
		zett.instances[key] = instance
	}
	return nil
}

func (zett *ZettContainer) IsBind(key string) bool{
	return zett.findServiceProvider != nil
}

func (zett *ZettContainer) Make(key string) (interface{},error){
	return zett.make(key,false,nil)
}
func (zett *ZettContainer) MustMake(key string) interface{}{
	instance,err := zett.make(key,false)
	if err != nil{
		panic(err)
	}
	return instance
}

func (zett *ZettContainer) MakeNew(key string,params ...interface{}) (interface{},error){
	return zett.make(key,true,params)
}


func (zett *ZettContainer) findServiceProvider(key string) ServiceProvider{
	zett.lock.RLock()
	defer zett.lock.RUnlock()
	if item,ok := zett.providers[key];ok{
		return item
	}
	return nil
}

func (zett *ZettContainer) make(key string,forceNew bool,params ...interface{}) (interface{},error){
	provider := zett.findServiceProvider(key)
	if provider == nil{
		return nil,errors.New(fmt.Sprintf("contract %s have no register",key))
	}
	if forceNew{
		return zett.newInstance(provider,params)
	}
	zett.lock.Lock()
	defer zett.lock.Unlock()
	if instance,ok:=zett.instances[key];ok{
		return instance,nil
	}
	instance,err := zett.newInstance(provider)
	if err != nil{
		return nil,err
	}
	zett.instances[key] = instance
	return instance,nil
}

func (zett *ZettContainer) newInstance(sp ServiceProvider,params ...interface{}) (interface{},error){
	if err := sp.Boot(zett);err != nil{
		return nil,err
	}
	if params == nil{
		params = sp.Params(zett)
	}
	method := sp.Register(zett)
	instance,err := method(params...)
	if err != nil{
		return nil,err
	}
	return instance,nil
}