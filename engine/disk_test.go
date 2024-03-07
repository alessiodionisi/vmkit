package engine

// func TestEngine_applyDisk(t *testing.T) {
// 	type fields struct {
// 		disks map[string]*Disk
// 	}

// 	type args struct {
// 		parsedDisk *v1beta1.Disk
// 	}

// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		want    bool
// 		wantErr bool
// 	}{
// 		{
// 			"create new disk",
// 			fields{
// 				disks: make(map[string]*Disk),
// 			},
// 			args{
// 				parsedDisk: &v1beta1.Disk{
// 					Metadata: &api.Metadata{
// 						Name: "test-disk",
// 					},
// 					Spec: &v1beta1.DiskSpec{
// 						Size: "5Gi",
// 					},
// 				},
// 			},
// 			true,
// 			false,
// 		},
// 		{
// 			"resize existing disk",
// 			fields{
// 				disks: map[string]*Disk{
// 					"test-disk": {
// 						Disk: &v1beta1.Disk{
// 							Metadata: &api.Metadata{
// 								Name: "test-disk",
// 							},
// 							Spec: &v1beta1.DiskSpec{
// 								Size: "5Gi",
// 							},
// 						},
// 					},
// 				},
// 			},
// 			args{
// 				parsedDisk: &v1beta1.Disk{
// 					Metadata: &api.Metadata{
// 						Name: "test-disk",
// 					},
// 					Spec: &v1beta1.DiskSpec{
// 						Size: "10Gi",
// 					},
// 				},
// 			},
// 			false,
// 			false,
// 		},
// 		{
// 			"change source of existing disk",
// 			fields{
// 				disks: map[string]*Disk{
// 					"test-disk": {
// 						Disk: &v1beta1.Disk{
// 							Metadata: &api.Metadata{
// 								Name: "test-disk",
// 							},
// 							Spec: &v1beta1.DiskSpec{
// 								Source: &v1beta1.DiskSpecSource{
// 									URL: "https://example.com/disk.qcow2",
// 								},
// 							},
// 						},
// 					},
// 				},
// 			},
// 			args{
// 				parsedDisk: &v1beta1.Disk{
// 					Metadata: &api.Metadata{
// 						Name: "test-disk",
// 					},
// 					Spec: &v1beta1.DiskSpec{
// 						Source: &v1beta1.DiskSpecSource{
// 							URL: "https://example.com/disk2.qcow2",
// 						},
// 					},
// 				},
// 			},
// 			false,
// 			true,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tmpDir, err := os.MkdirTemp("", "vmkit")
// 			if err != nil {
// 				t.Errorf("Engine.applyDisk() error creating temp dir: %s", err)
// 			}

// 			e, err := New(
// 				slog.Default(),
// 				tmpDir,
// 			)
// 			if err != nil {
// 				t.Errorf("Engine.applyDisk() error creating engine: %s", err)
// 			}

// 			e.disks = tt.fields.disks

// 			got, err := e.ApplyDisk(tt.args.parsedDisk)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("Engine.applyDisk() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}

// 			if got != tt.want {
// 				t.Errorf("Engine.applyDisk() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
