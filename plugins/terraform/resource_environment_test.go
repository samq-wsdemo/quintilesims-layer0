package main

import (
	"github.com/golang/mock/gomock"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/quintilesims/layer0/common/models"
	"testing"
)

func TestEnvironmentCreate_defaults(t *testing.T) {
	ctrl, mockClient, provider := setupUnitTest(t)
	defer ctrl.Finish()

	mockClient.EXPECT().
		CreateEnvironment("test-env", "m3.medium", 0, []byte(""), "linux", "").
		Return(&models.Environment{EnvironmentID: "eid"}, nil)

	mockClient.EXPECT().
		GetEnvironment("eid").
		Return(&models.Environment{}, nil)

	environmentResource := provider.ResourcesMap["layer0_environment"]
	d := schema.TestResourceDataRaw(t, environmentResource.Schema, map[string]interface{}{
		"name": "test-env",
	})

	if err := environmentResource.Create(d, mockClient); err != nil {
		t.Fatal(err)
	}
}

func TestEnvironmentCreate_specifyOptional(t *testing.T) {
	ctrl, mockClient, provider := setupUnitTest(t)
	defer ctrl.Finish()

	mockClient.EXPECT().
		CreateEnvironment("test-env", "m3.large", 2, []byte("user data"), "windows", "ami_id").
		Return(&models.Environment{EnvironmentID: "eid"}, nil)

	mockClient.EXPECT().
		GetEnvironment("eid").
		Return(&models.Environment{}, nil)

	environmentResource := provider.ResourcesMap["layer0_environment"]
	d := schema.TestResourceDataRaw(t, environmentResource.Schema, map[string]interface{}{
		"name":      "test-env",
		"size":      "m3.large",
		"min_count": 2,
		"user_data": "user data",
		"os":        "windows",
		"ami":       "ami_id",
	})

	if err := environmentResource.Create(d, mockClient); err != nil {
		t.Fatal(err)
	}
}

func TestEnvironmentRead(t *testing.T) {
	ctrl, mockClient, provider := setupUnitTest(t)
	defer ctrl.Finish()

	mockClient.EXPECT().
		GetEnvironment("eid").
		Return(&models.Environment{}, nil)

	environmentResource := provider.ResourcesMap["layer0_environment"]
	d := schema.TestResourceDataRaw(t, environmentResource.Schema, map[string]interface{}{})
	d.SetId("eid")

	if err := environmentResource.Read(d, mockClient); err != nil {
		t.Fatal(err)
	}
}

func TestEnvironmentUpdate(t *testing.T) {
	ctrl, mockClient, provider := setupUnitTest(t)
	defer ctrl.Finish()

	gomock.InOrder(
		mockClient.EXPECT().
			CreateEnvironment("test-env", "m3.medium", 0, []byte(""), "linux", "").
			Return(&models.Environment{EnvironmentID: "eid"}, nil),

		mockClient.EXPECT().
			GetEnvironment("eid").
			Return(&models.Environment{EnvironmentID: "eid"}, nil),

		mockClient.EXPECT().
			UpdateEnvironment("eid", 3).
			Return(&models.Environment{EnvironmentID: "eid"}, nil),

		mockClient.EXPECT().
			GetEnvironment("eid").
			Return(&models.Environment{EnvironmentID: "eid"}, nil),
	)

	environmentResource := provider.ResourcesMap["layer0_environment"]
	d1 := schema.TestResourceDataRaw(t, environmentResource.Schema, map[string]interface{}{
		"name": "test-env",
	})

	d2 := schema.TestResourceDataRaw(t, environmentResource.Schema, map[string]interface{}{
		"name":      "test-env",
		"min_count": 3,
	})

	d2.SetId("eid")

	if err := environmentResource.Create(d1, mockClient); err != nil {
		t.Fatal(err)
	}

	if err := environmentResource.Update(d2, mockClient); err != nil {
		t.Fatal(err)
	}
}

func TestEnvironmentDelete(t *testing.T) {
	ctrl, mockClient, provider := setupUnitTest(t)
	defer ctrl.Finish()

	mockClient.EXPECT().
		DeleteEnvironment("eid").
		Return("jid", nil)

	mockClient.EXPECT().
		WaitForJob("jid", gomock.Any()).
		Return(nil)

	environmentResource := provider.ResourcesMap["layer0_environment"]
	d := schema.TestResourceDataRaw(t, environmentResource.Schema, map[string]interface{}{})
	d.SetId("eid")

	if err := environmentResource.Delete(d, mockClient); err != nil {
		t.Fatal(err)
	}
}