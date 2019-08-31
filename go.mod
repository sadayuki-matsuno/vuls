module github.com/sadayuki-matsuno/vuls

go 1.12

require (
	github.com/Azure/go-autorest/autorest/to v0.3.0 // indirect
	github.com/Sirupsen/logrus v1.4.0 // indirect
	github.com/elazarl/goproxy v0.0.0-20190711103511-473e67f1d7d2 // indirect
	github.com/elazarl/goproxy/ext v0.0.0-20190711103511-473e67f1d7d2 // indirect
	github.com/google/subcommands v1.0.1
	github.com/howeyc/gopass v0.0.0-20170109162249-bf9dde6d0d2c
	github.com/k0kubun/pp v3.0.1+incompatible
	github.com/kotakanbe/go-cve-dictionary v0.3.1
	github.com/kotakanbe/go-pingscanner v0.1.0
	github.com/mitchellh/go-homedir v1.1.0
	github.com/nsf/termbox-go v0.0.0-20190817171036-93860e161317 // indirect
	github.com/onsi/ginkgo v1.10.1 // indirect
	github.com/onsi/gomega v1.7.0 // indirect
	github.com/sadayuki-matsuno/vuls-config v0.0.0-20190831051520-f427ddca9f2c
	github.com/sadayuki-matsuno/vuls-models v0.0.0-20190831051724-bd249f9152cc
	github.com/sadayuki-matsuno/vuls-report v0.0.0-20190831055402-14fe5b1ca5cc
	github.com/sadayuki-matsuno/vuls-scan v0.0.0-20190831055612-40ddc48e8f9b
	github.com/sirupsen/logrus v1.4.2
	github.com/smartystreets/goconvey v0.0.0-20190731233626-505e41936337 // indirect
	golang.org/x/xerrors v0.0.0-20190717185122-a985d3407aa7
	gopkg.in/airbrake/gobrake.v2 v2.0.9 // indirect
	gopkg.in/gemnasium/logrus-airbrake-hook.v2 v2.1.2 // indirect
)

replace (
	github.com/Sirupsen/logrus v1.0.5 => github.com/sirupsen/logrus v1.0.5

	github.com/Sirupsen/logrus v1.3.0 => github.com/Sirupsen/logrus v1.0.6
	github.com/Sirupsen/logrus v1.4.0 => github.com/sirupsen/logrus v1.0.6
	github.com/genuinetools/reg => github.com/tomoyamachi/reg v0.16.1-0.20190706172545-2a2250fd7c00
	gopkg.in/mattn/go-colorable.v0 => github.com/mattn/go-colorable v0.1.0
	gopkg.in/mattn/go-isatty.v0 => github.com/mattn/go-isatty v0.0.6
)
