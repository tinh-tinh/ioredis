package throttler

import "github.com/tinh-tinh/tinhtinh/core"

func Guard(name string) core.Guard {
	return func(module *core.DynamicModule, ctx core.Ctx) bool {
		throttler := module.Ref(core.Provide(name)).(*Throttler)
		ip := ctx.Headers("X-Real-Ip")
		if ip == "" {
			ip = ctx.Headers("X-Forwarded-For")
		}

		hits := throttler.Get(ip)
		if hits > throttler.Max {
			return false
		}
		throttler.Incr(ip)

		return true
	}
}
