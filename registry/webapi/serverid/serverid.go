package serverid

import (
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"

	"webapi/details"
)

const (
	colonDelimiter   = ":"
	randomnessLength = 16
)

var (
	serverID, errServerID = createServerID(randomnessLength)
)

func generateRandomByteArray(n int, err error) (*[]byte, error) {
	if err != nil {
		return nil, err
	}

	token := make([]byte, n)
	length, errRandom := rand.Read(token)
	if errRandom != nil || length != n {
		return nil, errRandom
	}

	return &token, nil
}

func createServerID(randomnessLength int) (*string, error) {
	serviceName := details.ConfDetails.ServiceName
	now := time.Now().Unix()
	randomBytes, errRandomBytes := generateRandomByteArray(randomnessLength, nil)
	if errRandomBytes != nil {
		return nil, errRandomBytes
	}

	bytesAsStr := hex.EncodeToString(*randomBytes)

	preparedID := fmt.Sprint(
		now,
		colonDelimiter,
		serviceName,
		colonDelimiter,
		bytesAsStr,
	)

	return &preparedID, nil
}
