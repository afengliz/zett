package contract

import "net/http"

const KernelKey = "zett:kernel"

type Kernel interface {
	HttpEngine() http.Handler
}
