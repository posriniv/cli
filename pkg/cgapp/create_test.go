package cgapp

import (
	"os"
	"testing"
)

func TestCreateProjectFromRegistry(t *testing.T) {
	type args struct {
		p        *Project
		registry map[string]string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"successfully create default Ansible roles",
			args{
				p: &Project{
					Name:       "roles",
					Type:       "roles",
					RootFolder: "../../tmp",
				},
				registry: map[string]string{
					"roles": "github.com/create-go-app/ansible-roles",
				},
			},
			false,
		},
		{
			"successfully create default backend",
			args{
				p: &Project{
					Name:       "echo",
					Type:       "backend",
					RootFolder: "../../tmp",
				},
				registry: map[string]string{
					"echo": "github.com/create-go-app/echo-go-template",
				},
			},
			false,
		},
		{
			"successfully create default webserver",
			args{
				p: &Project{
					Name:       "nginx",
					Type:       "webserver",
					RootFolder: "../../tmp",
				},
				registry: map[string]string{
					"nginx": "github.com/create-go-app/nginx-docker",
				},
			},
			false,
		},
		{
			"successfully create default database",
			args{
				p: &Project{
					Name:       "postgres",
					Type:       "database",
					RootFolder: "../../tmp",
				},
				registry: map[string]string{
					"postgres": "github.com/create-go-app/postgres-docker",
				},
			},
			false,
		},
		{
			"successfully create backend from user template",
			args{
				p: &Project{
					Name:       "github.com/create-go-app/echo-go-template",
					Type:       "backend",
					RootFolder: "../../tmp",
				},
				registry: map[string]string{},
			},
			false,
		},
		{
			"failed create default database",
			args{
				p: &Project{
					Name:       "",
					Type:       "",
					RootFolder: "../../tmp",
				},
				registry: map[string]string{},
			},
			true,
		},
		{
			"failed create default database",
			args{
				p: &Project{
					Name:       "",
					Type:       "fff",
					RootFolder: "../../tmp",
				},
				registry: map[string]string{},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateProjectFromRegistry(tt.args.p, tt.args.registry); (err != nil) != tt.wantErr {
				t.Errorf("CreateProjectFromRegistry() error = %v, wantErr %v", err, tt.wantErr)
			}

			// Clean
			os.RemoveAll(tt.args.p.RootFolder)
		})
	}
}
