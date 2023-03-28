// Copyright (C) 2022-2023	 Akira Tanimura (@autopp)
//
// Licensed under the Apache License, Version 2.0 (the “License”);
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an “AS IS” BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/docker/docker/pkg/fileutils"
	"github.com/moby/buildkit/frontend/dockerfile/dockerignore"
	"github.com/spf13/cobra"
)

func Main(version string, stdin io.Reader, stdout, stderr io.Writer, args []string) error {
	versionFlag := "version"

	cmd := &cobra.Command{
		Use:           "digrep [DIR]",
		SilenceErrors: true,
		SilenceUsage:  true,
		Args:          cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if v, err := cmd.Flags().GetBool(versionFlag); err != nil {
				return err
			} else if v {
				cmd.Println(version)
				return nil
			}

			var patterns []string

			ignorefile := ".dockerignore"
			if len(args) > 0 {
				ignorefile = filepath.Join(args[0], ignorefile)
			}
			f, err := os.Open(ignorefile)

			if err == nil {
				defer f.Close()
				patterns, err = dockerignore.ReadAll(f)
				if err != nil {
					fmt.Fprintf(stderr, "cannot read .dockerignore: %s", err)
					return err
				}
			} else if !os.IsNotExist(err) {
				fmt.Fprintf(stderr, "cannot open .dockerignore: %s", err)
				return err
			}

			entries := make([]string, 0)
			s := bufio.NewScanner(stdin)
			for s.Scan() {
				entries = append(entries, s.Text())
			}

			if err := s.Err(); err != nil {
				fmt.Fprintf(stderr, "cannot read stdin: %s", err)
			}

			pm, err := fileutils.NewPatternMatcher(patterns)
			if err != nil {
				fmt.Fprintf(stderr, "invalid pattern: %s", err)
				return err
			}

			for _, entry := range entries {
				ignored, err := pm.Matches(entry)
				if err != nil {
					fmt.Fprintf(stderr, "cannot match %s: %s", entry, err)
					return err
				}
				if !ignored {
					fmt.Println(entry)
				}
			}

			return nil
		},
	}

	cmd.Flags().Bool(versionFlag, false, "print version")

	cmd.SetIn(stdin)
	cmd.SetOut(stdout)
	cmd.SetErr(stderr)
	cmd.SetArgs(args)

	return cmd.Execute()
}
