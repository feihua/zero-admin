package shared

import "fmt"

var (
	VoteUidAuidLockKey = "vote.lock#%v#%v#"
	UidAuidLockKey     = "user.lock#%v#%v#"
)

func GetUidAuidLockKey(uid, auid int64) string {
	return fmt.Sprintf(UidAuidLockKey, uid, auid)
}

func GetVoteUidAuidLockKey(uid, auid int64) string {
	return fmt.Sprintf(VoteUidAuidLockKey, uid, auid)
}
