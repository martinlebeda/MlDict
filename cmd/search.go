// Copyright © 2018 Martin Lebeda <martin.lebeda@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"github.com/martinlebeda/mldict/service"
	"github.com/martinlebeda/mldict/termout"
	"github.com/spf13/cobra"
)

var dict string
var exact bool

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:     "search",
	Aliases: []string{"find", "s"},
	Args:    cobra.ExactArgs(1),
	Short:   "Search dictionaries for a word",
	Long: `Search dictionary database for a word. 

For searching use fulltext index in sqlite3 (FT4). 

Implicit is search for "term*", it is for key starting with term. 
You can owerride this by '-e' option and search for exact term.`,
	Run: func(cmd *cobra.Command, args []string) {
		termout.PrintResult(service.QueryDict(args[0], exact, dict))
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

	searchCmd.Flags().BoolVarP(&exact, "exact", "e", false, "Search for exact term")
	searchCmd.Flags().StringVarP(&dict, "dict", "d", "", "Select one dictionary")
}
