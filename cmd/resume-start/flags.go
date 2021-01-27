package main

import "flag"

const deafultResumeGitDirectory = "./resumake.io"
const defaultLogFile = "./.resume-start.log"
const defaultSkipBuild = false
const defaultNoClient = false

type cliFlags struct {
	resumeGitDirectory string
	logFile            string
	skipBuild          bool
	noClient           bool
}

func getFlags() cliFlags {
	f := cliFlags{}
	flag.StringVar(&f.resumeGitDirectory, "resumake-dir", deafultResumeGitDirectory, "the directory where resumake.io resides")
	flag.StringVar(&f.logFile, "log", defaultLogFile, "the file where resume-start logs to")
	flag.BoolVar(&f.skipBuild, "skip-build", defaultSkipBuild, "skip building resumake.io dependencies")
	flag.BoolVar(&f.noClient, "no-client", defaultNoClient, "disable running the client")
	flag.Parse()
	return f
}
