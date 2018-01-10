/*
 * gocmd
 * For the full copyright and license information, please view the LICENSE.txt file.
 */

package flagset

// Command represents a command
type Command struct {
	id        int
	command   string
	flagID    int
	parentID  int
	argID     int
	indexFrom int
	indexTo   int
	updatedBy []string // for debug
	err       error
}
