package main

import "flag"

const deafultResumeGitDirectory = "./resumake.io"
const defaultLogFile = ".resume-start.log"
const defaultSkipBuild = false
const defaultServerOnly = false

type cliFlags struct {
	resumeGitDirectory string
	logFile            string
	skipBuild          bool
	serverOnly         bool
}

func getFlags() cliFlags {
	f := cliFlags{}
	flag.StringVar(&f.resumeGitDirectory, "resumake-dir", deafultResumeGitDirectory, "the directory where resumake.io resides")
	flag.StringVar(&f.logFile, "log", defaultLogFile, "the file where resume-start logs too")
	flag.BoolVar(&f.skipBuild, "skip-build", defaultSkipBuild, "skip building resumake.io dependencies")
	flag.BoolVar(&f.serverOnly, "server-only", defaultServerOnly, "only run the resumake.io's server")
	flag.Parse()
	return f
}
