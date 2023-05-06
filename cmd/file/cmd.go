package file

import "github.com/urfave/cli/v2"

var Command = cli.Command{
	Name: "file",
	Subcommands: []*cli.Command{
		{
			Name:   "list",
			Action: cmdList,
			Flags: []cli.Flag{
				flagFileID,
			},
		},
		{
			Name:   "upload",
			Action: cmdUpload,
			Flags: []cli.Flag{
				rflagFile,
				rflagPurpose,
			},
		},
		{
			Name:   "delete",
			Action: cmdDelete,
			Flags: []cli.Flag{
				rflagFileID,
			},
		},
		{
			Name:   "content",
			Action: cmdContent,
			Flags: []cli.Flag{
				rflagFileID,
			},
		},
	},
}

var flagFileID = &cli.StringFlag{
	Name: "file_id",
}

var rflagFileID = &cli.StringFlag{
	Name:     "file_id",
	Required: true,
}

var rflagFile = &cli.StringFlag{
	Name:     "file",
	Required: true,
}

var rflagPurpose = &cli.StringFlag{
	Name:     "purpose",
	Required: true,
}
