package base

import "github.com/urfave/cli/v2"

func Text1Flag() *cli.StringFlag {
	f := cli.StringFlag {     
	  Name: "text1",
	  Aliases: []string{"t1"},
	  Usage: "Show Name",
	}
	return &f	
}

func Text2Flag() *cli.StringFlag {
	f := cli.StringFlag {     
	  Name: "text2",
	  Aliases: []string{"t2"},
	  Usage: "Show Name",
	}
	return &f	
}

func Arg1Flag() *cli.StringFlag {
	f := cli.StringFlag {     
	  Name: "arg1",
	  Aliases: []string{"a1"},
	  Usage: "Args1",
	}
	return &f
}

func Arg2Flag() *cli.StringFlag {
	f := cli.StringFlag {     
	  Name: "arg2",
	  Aliases: []string{"a2"},
	  Usage: "Args2",
	}
	return &f
}

func HostFlag() *cli.StringFlag {
	 f := cli.StringFlag {     
	   Name: "host",
	   Aliases: []string{"ho"},
	   Usage: "Host Flag",
	 }
	 return &f 
}


func LoadFlag() *cli.StringFlag {
	f := cli.StringFlag {     
	  Name: "yaml",
	  Aliases: []string{"yml"},
	  Usage: "Load Yaml/Yml Flag",
	}
	return &f 
}