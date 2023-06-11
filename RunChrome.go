package main

import (
	"os"
	"fmt"
	"context"
   "github.com/chromedp/chromedp"
)

func RunChrome(userDir, listenPort, address string) (context.Context, context.CancelFunc) {
	var ctx context.Context
	var cancel context.CancelFunc
	var opts []chromedp.ExecAllocatorOption
	var wd string
	var err error

	Goose.Init.Logf(2, "Starting Chrome ...")
	defer Goose.Init.Logf(2, "Chrome started ...")

/*
	opts = append(chromedp.DefaultExecAllocatorOptions[:],
//		chromedp.ProxyServer("http://127.0.0.1:8118"), 
//		chromedp.ProxyServer("socks5://127.0.0.1:9050"),
//		chromedp.Flag("proxy-server", ";socks=127.0.0.1:9050;sock4=127.0.0.1:9050;sock5=127.0.0.1:9050"),
//		chromedp.Flag("proxy-server", "http://127.0.0.1:8118"),
		chromedp.NoSandbox,
		chromedp.Headless,
		chromedp.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.163 Safari/537.36"),
		chromedp.WindowSize(WindowWidth, WindowHeight),
//		chromedp.Flag("remote-debugging-port", 9222),
	)

	if useProxy {
		opts = append(opts,
//			  chromedp.ProxyServer("http://127.0.0.1:8118"), 
			chromedp.ProxyServer("socks5://127.0.0.1:9050"),
//			chromedp.Flag("proxy-server", ";socks=127.0.0.1:9050;sock4=127.0.0.1:9050;sock5=127.0.0.1:9050"),
//			chromedp.Flag("proxy-server", "http://127.0.0.1:8118"),
		)
	}

*/

	wd, err = os.Getwd()
	if err != nil {
		Goose.Init.Fatalf(2, "Chrome start failed...")
	}


	opts = append(chromedp.DefaultExecAllocatorOptions[:],

//		chromedp.NoSandbox,
		chromedp.NoFirstRun,
		chromedp.NoDefaultBrowserCheck,
//		chromedp.DisableGPU,

////		chromedp.Flag("remote-debugging-port", 9222),

		// puppeteer default behavior
		chromedp.Flag("disable-infobars", true),
		//// chromedp.Flag("excludeSwitches", "enable-automation"),
		chromedp.Flag("disable-background-networking", true),
		chromedp.Flag("enable-features", "NetworkService,NetworkServiceInProcess"),
		chromedp.Flag("disable-background-timer-throttling", true),
		chromedp.Flag("disable-backgrounding-occluded-windows", true),
		chromedp.Flag("disable-breakpad", true),
		chromedp.Flag("enable-lazy-image-loading", false),
		chromedp.Flag("disable-client-side-phishing-detection", true),
		chromedp.Flag("disable-default-apps", true),
		chromedp.Flag("disable-dev-shm-usage", true),
		chromedp.Flag("disable-extensions", true),
		chromedp.Flag("disable-features", "site-per-process,TranslateUI,BlinkGenPropertyTrees"),
		chromedp.Flag("disable-ipc-flooding-protection", true),
		chromedp.Flag("disable-popup-blocking", true),
		chromedp.Flag("disable-prompt-on-repost", true),
		chromedp.Flag("disable-renderer-backgrounding", true),
		chromedp.Flag("disable-sync", true),
		chromedp.Flag("force-color-profile", "srgb"),
		chromedp.Flag("metrics-recording-only", true),
		chromedp.Flag("safebrowsing-disable-auto-update", true),
		chromedp.Flag("password-store", "basic"),
		chromedp.Flag("use-mock-keychain", true),
//		chromedp.Flag("use-fake-ui-for-media-stream", true),
//		chromedp.Flag("use-fake-device-for-media-stream", true),

		// custom args
		chromedp.Flag("incognito", true),
		chromedp.Flag("disable-prompt-for-download", true),
//		chromedp.Flag("kiosk", true),
		chromedp.Flag("kiosk", false),
		chromedp.Flag("app", "data:,"),
/////////////////////////////////////////////////////

		chromedp.Flag("disable-features", "StrictOriginIsolation"),
		chromedp.Flag("change-stack-guard-on-fork", "enable"),
		chromedp.Flag("no-periodic-tasks", true),

//////////////

		chromedp.Flag("disable-hang-monitor", true),

		chromedp.Flag("headless", false),
//		chromedp.Headless,

		// disable security
//		chromedp.Flag("disable-web-security", true),
//// 		chromedp.Flag("allow-running-insecure-content", true),
//		chromedp.Flag("ignore-certificate-errors", true),

//		chromedp.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.163 Safari/537.36"),
		chromedp.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.181 Safari/537.36 WAIT_UNTIL=load"),
		chromedp.WindowSize(WindowWidth, WindowHeight),
		chromedp.Flag("window-position", "0,0"),
		chromedp.Flag("window-size", fmt.Sprintf("%d,%d", WindowWidth, WindowHeight)),
		chromedp.Flag("autoplay-policy", "no-user-gesture-required"),
		chromedp.Flag("enable-automation", false),
//		chromedp.Flag("enable-automation", true),
		chromedp.Flag("user-data-dir", fmt.Sprintf("%s%c%s", wd, os.PathSeparator, userDir)),
		chromedp.Flag("profile-directory", "Default"),
		chromedp.Flag("proxy-server", "http://127.0.0.1:" + listenPort),
		chromedp.Flag("host-resolver-rules", "MAP * ~NOTFOUND , EXCLUDE 127.0.0.1"),
	)

	ctx, _ = chromedp.NewExecAllocator(context.Background(), opts...)
	ctx, cancel = chromedp.NewContext(ctx)

//	ctx, cancel = context.WithTimeout(ctx, 25*time.Second)

	_ = chromedp.Run(ctx, chromedp.Navigate(address))

	return ctx, cancel
}
