package door

import (
	"errors"
	"testing"
)

type TestScenario struct {
	name           string
	args           interface{}
	door           Door
	expectedResult interface{}
	expectedError  error
	expectedValue  interface{}
}

func TestNewDoor(t *testing.T) {
	testScenarios := []TestScenario{
		{
			name:          "Should create new Door",
			args:          "ValidKey",
			expectedError: nil,
		},
		{
			name:          "Should return error",
			args:          "",
			expectedError: ErrInvalidKey,
		},
	}

	for _, test := range testScenarios {
		t.Run(test.name, func(t *testing.T) {
			key := test.args.(string)
			door, err := NewDoor(key)
			if door == test.expectedResult {
				t.Errorf("NewDoor() returns %#v but expected %#v", door, test.expectedResult)
			}

			if !errors.Is(err, test.expectedError) {
				t.Errorf("NewDoor() returns errror %#v but expected %#v", err, test.expectedError)
			}
		})
	}
}

func TestDoorOpen(t *testing.T) {
	testScenarios := []TestScenario{
		{
			name:           "Closed door should be opened",
			door:           Door{isClosed: true},
			expectedResult: nil,
			expectedValue:  true,
		},
		{
			name:          "Open door should return error",
			door:          Door{isClosed: false},
			expectedError: ErrOpenDoor,
			expectedValue: true,
		},
		{
			name:          "Locked door should return error",
			door:          Door{isLocked: true, isClosed: true},
			expectedError: ErrLockedDoor,
			expectedValue: false,
		},
	}

	for _, test := range testScenarios {
		t.Run(test.name, func(t *testing.T) {
			door := test.door
			result := door.Open()
			if (test.expectedResult != nil && result != test.expectedResult) || (test.expectedError != nil && result != test.expectedError) {
				t.Errorf("Open() returns %#v but expected %#v", result, test.expectedResult)
			}

			if door.IsOpen() != test.expectedValue {
				t.Errorf("door.IsOpen is %#v but expected %#v", door.IsOpen(), test.expectedValue)
			}
		})
	}
}

func TestDoorClose(t *testing.T) {
	testScenarios := []TestScenario{
		{
			name:           "Open door should be closed",
			door:           Door{isClosed: false},
			expectedResult: nil,
			expectedValue:  true,
		},
		{
			name:          "Closed door should return error",
			door:          Door{isClosed: true},
			expectedError: ErrClosedDoor,
			expectedValue: true,
		},
	}

	for _, test := range testScenarios {
		t.Run(test.name, func(t *testing.T) {
			door := test.door
			result := door.Close()
			if (test.expectedResult != nil && result != test.expectedResult) || (test.expectedError != nil && result != test.expectedError) {
				t.Errorf("Close() returns %#v but expected %#v", result, test.expectedResult)
			}

			if door.IsClosed() != test.expectedValue {
				t.Errorf("door.isClosed is %#v but expected %#v", door.IsClosed(), test.expectedValue)
			}
		})
	}
}

func TestDoorLockShouldLockedDoor(t *testing.T) {
	door, err := NewDoor("ValidKey")
	if err != nil {
		t.Errorf("TestDoorLock returns error %#v", err)
	}

	door.Lock()

	if !door.IsLocked() {
		t.Errorf("door.isLocked is %#v but expected %#v", door.IsLocked(), true)
	}
}

func TestDoorUnlock(t *testing.T) {
	testScenarios := []TestScenario{
		{
			name:           "Locked door should be unlocked",
			door:           Door{key: "ValidKey", isLocked: true},
			args:           "ValidKey",
			expectedResult: nil,
			expectedValue:  true,
			expectedError:  nil,
		},
		{
			name:           "Locked door with wrong key should return error",
			door:           Door{key: "ValidKey", isLocked: true},
			args:           "InvalidKey",
			expectedResult: nil,
			expectedError:  ErrWrongKey,
			expectedValue:  false,
		},
	}

	for _, test := range testScenarios {
		t.Run(test.name, func(t *testing.T) {
			door := test.door
			key := test.args.(string)
			result := door.Unlock(key)
			if (test.expectedResult != nil && result != test.expectedResult) || (test.expectedError != nil && result != test.expectedError) {
				t.Errorf("Unlock() returns %#v but expected %#v", result, test.expectedResult)
			}

			if door.IsUnlocked() != test.expectedValue {
				t.Errorf("Door.IsUnlocked is %#v but expected %#v", door.IsUnlocked(), test.expectedValue)
			}
		})
	}
}
