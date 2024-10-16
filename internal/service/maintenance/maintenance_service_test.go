// Copyright © 2022 FuryLion Group LLC. All Rights Reserved.
//
// Unauthorized copying of this file, via any medium is strictly prohibited
// without the express permission of FuryLion Group LLC.
// Proprietary and confidential.
//
// Created by Vadim Vlasov on 11.10.22

package maintenance

import (
	"reflect"
	"testing"
	"time"

	"locgame-mini-server/pkg/dto/base"
	"locgame-mini-server/pkg/dto/maintenance"
)

func Test_findCloserMaintenance(t *testing.T) {
	type args struct {
		data []*maintenance.MaintenanceData
		now  int64
	}
	tests := []struct {
		name string
		args args
		want *maintenance.MaintenanceData
	}{
		{name: "Maintenance start date has not yet arrived", args: struct {
			data []*maintenance.MaintenanceData
			now  int64
		}{
			data: []*maintenance.MaintenanceData{
				{
					StartDate: &base.Timestamp{
						Seconds: time.Date(2022, 11, 2, 0, 0, 0, 0, time.UTC).Unix(),
					},
					EndDate: &base.Timestamp{
						Seconds: time.Date(2022, 11, 2, 2, 0, 0, 0, time.UTC).Unix(),
					},
				},
				{
					StartDate: &base.Timestamp{
						Seconds: time.Date(2022, 12, 3, 0, 0, 0, 0, time.UTC).Unix(),
					},
					EndDate: &base.Timestamp{
						Seconds: time.Date(2022, 12, 3, 2, 0, 0, 0, time.UTC).Unix(),
					},
				},
				{
					StartDate: &base.Timestamp{
						Seconds: time.Date(2022, 12, 2, 0, 0, 0, 0, time.UTC).Unix(),
					},
					EndDate: &base.Timestamp{
						Seconds: time.Date(2022, 12, 2, 2, 0, 0, 0, time.UTC).Unix(),
					},
				},
				{
					StartDate: &base.Timestamp{
						Seconds: time.Date(2022, 12, 2, 3, 0, 0, 0, time.UTC).Unix(),
					},
					EndDate: &base.Timestamp{
						Seconds: time.Date(2022, 12, 2, 4, 0, 0, 0, time.UTC).Unix(),
					},
				},
			},
			now: time.Date(2022, 12, 1, 0, 0, 0, 0, time.UTC).Unix()},
			want: &maintenance.MaintenanceData{
				StartDate: &base.Timestamp{
					Seconds: time.Date(2022, 12, 2, 0, 0, 0, 0, time.UTC).Unix(),
				},
				EndDate: &base.Timestamp{
					Seconds: time.Date(2022, 12, 2, 2, 0, 0, 0, time.UTC).Unix(),
				},
			},
		},

		{
			name: "Maintenance data is empty",
			args: struct {
				data []*maintenance.MaintenanceData
				now  int64
			}{
				now:  time.Date(2022, 12, 1, 0, 0, 0, 0, time.UTC).Unix(),
				data: []*maintenance.MaintenanceData{},
			},
			want: nil,
		},
	}
	s := new(Service)
	for _, tt := range tests {
		testData := tt
		t.Run(testData.name, func(t *testing.T) {
			if got := s.findCloserMaintenance(testData.args.data, testData.args.now); !reflect.DeepEqual(got, testData.want) {
				t.Errorf("findCloserMaintenance() = %v, want %v", got, testData.want)
			}
		})
	}
}
