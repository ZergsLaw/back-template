// Code generated by "stringer -output=stringer.LogKey.go -type=LogKey -linecomment"; DO NOT EDIT.

package logger

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Version-1]
	_ = x[PanicReason-2]
	_ = x[URL-3]
	_ = x[Error-4]
	_ = x[Reason-5]
	_ = x[Host-6]
	_ = x[Port-7]
	_ = x[Module-8]
	_ = x[Environment-9]
	_ = x[Stack-10]
	_ = x[TaskID-11]
	_ = x[TaskKind-12]
}

const _LogKey_name = "versionpanic_reasonurlerrreasonhostportmoduleenvironmentstacktask_idtask_kind"

var _LogKey_index = [...]uint8{0, 7, 19, 22, 25, 31, 35, 39, 45, 56, 61, 68, 77}

func (i LogKey) String() string {
	i -= 1
	if i >= LogKey(len(_LogKey_index)-1) {
		return "LogKey(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _LogKey_name[_LogKey_index[i]:_LogKey_index[i+1]]
}
