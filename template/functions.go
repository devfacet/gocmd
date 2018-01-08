/*
 * gocmd
 * For the full copyright and license information, please view the LICENSE.txt file.
 */

package template

import (
	"os"
	"os/exec"
	"strings"
	"time"
)

// tplFuncEnv returns an environment variable
func tplFuncEnv(args ...interface{}) string {
	// Init vars
	result := ""

	if len(args) == 0 {
		return result
	}

	// Env key
	if k, ok := args[0].(string); ok && k != "" {
		if v, ok := os.LookupEnv(k); ok && v != "" {
			result = v
		}
	}

	// Default value
	if result == "" && len(args) > 1 {
		if s, ok := args[1].(string); ok && s != "" {
			result = s
		}
	}

	return result
}

// tplFuncTime returns a time by the given arguments
func tplFuncTime(args ...interface{}) string {
	// Init vars
	format := time.RFC3339 // profile of ISO8601 (2006-01-02T15:04:05Z07:00)
	t := time.Now()

	if len(args) == 0 {
		return t.Format(format)
	}

	// Parse options (i.e. foo=bar&hello=world)
	if opts, ok := args[0].(string); ok && opts != "" {
		l := strings.Split(opts, "&")
		if len(l) > 0 {
			for _, v := range l {
				kv := strings.Split(v, "=")
				switch kv[0] {
				case "format":
					if len(kv) > 1 {
						format = kv[1]
					}
				case "add":
					if len(kv) > 1 {
						if d, err := time.ParseDuration(kv[1]); err == nil {
							t = t.Add(d)
						}
					}
				case "sub":
					if len(kv) > 1 {
						if d, err := time.ParseDuration(kv[1]); err == nil {
							t = t.Add(-d)
						}
					}
				}
			}
		}
	}

	return t.Format(format)
}

// tplFuncExec executes a command by the given arguments
func tplFuncExec(args ...interface{}) string {
	// Init vars
	result := ""

	if len(args) == 0 {
		return result
	}

	// Command
	if c, ok := args[0].(string); ok && c != "" {
		// Build arguments
		ca := []string{}
		for _, v := range args[1:] {
			ca = append(ca, v.(string))
		}
		// Execute the command
		out, err := exec.Command(c, ca...).Output()
		if err == nil {
			result = strings.TrimSpace(string(out))
		}
	}

	return result
}
