package test_service

import (
	"context"
	"github.com/SameeranB/go-heavy/pkg"
	"log"
	"time"
)

type CustomParams struct {
	Message string
}

// Todo: Load this from a config file
type CustomConfig struct {
	BaseURL string
}

func main() pkg.Workflow {

	var setupFunc = func(pctx pkg.PassedContext) (pkg.PassedContext, error) {
		ctx, cancel := context.WithTimeout(pctx.Ctx, 5*time.Second)
		pctx.Ctx = ctx
		pctx.CancelFunc = cancel

		pctx.Config = CustomConfig{
			BaseURL: "httpbin.org",
		}

		pctx.Providers.UseHTTPProvider()

		log.Println("Staring Ping Test")
		return pctx, nil
	}

	var initFunc = func(pctx pkg.PassedContext) (pkg.PassedContext, error) {
		log.Printf("Running iteration %d. With user %s\n", pctx.Iter, pctx.UserId)

		pctx.Params = CustomParams{
			Message: "Hello World",
		}
		return pctx, nil
	}

	var pingStep = func(pctx pkg.PassedContext) (pkg.PassedContext, error) {
		log.Printf("Params Say: %s", pctx.Params.(CustomParams).Message)
		res, err := pctx.Providers.HTTPProvider.Get(
			pctx.Ctx,
			pctx.Config.(CustomConfig).BaseURL,
			"/get",
			nil,
			nil,
		)
		if err != nil {
			return pctx, err
		}
		log.Println(res)
		// Todo: Need an assertion package
		return pctx, nil
	}

	var steps = []func(pctx pkg.PassedContext) (pkg.PassedContext, error){
		pingStep,
	}

	var teardownFunc = func(pctx pkg.PassedContext) (pkg.PassedContext, error) {
		log.Println("Teardown")
		return pctx, nil
	}

	return pkg.Workflow{
		Name:     "Ping Workflow",
		SetUp:    setupFunc,
		Init:     initFunc,
		Steps:    steps,
		Teardown: teardownFunc,
	}
}
