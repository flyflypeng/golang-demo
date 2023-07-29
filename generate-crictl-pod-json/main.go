package main

import (
	"encoding/json"
	"fmt"
	"os"

	pb "k8s.io/cri-api/pkg/apis/runtime/v1"
)

func main() {
	podSandboxConfig := pb.PodSandboxConfig{
		Metadata: &pb.PodSandboxMetadata{
			Name:      "defaultName",
			Uid:       "defaultUid",
			Namespace: "defaultNamespace",
			Attempt:   1,
		},
		Hostname:     "defaultHostname",
		LogDirectory: "/var/log/pods/defaultNamespace_defaultName_defaultUid/",
		DnsConfig: &pb.DNSConfig{
			Servers:  []string{"8.8.8.8"},
			Searches: []string{"example.com"},
			Options:  []string{"ndots:5"},
		},
		PortMappings: []*pb.PortMapping{
			{
				Protocol:      pb.Protocol_TCP,
				ContainerPort: 8080,
				HostPort:      8080,
			},
		},
		Labels: map[string]string{
			"label1": "value1",
			"label2": "value2",
		},
		Annotations: map[string]string{
			"annotation1": "value1",
			"annotation2": "value2",
		},
		Linux: &pb.LinuxPodSandboxConfig{
			CgroupParent: "/docker/defaultCgroup",
			SecurityContext: &pb.LinuxSandboxSecurityContext{
				NamespaceOptions: &pb.NamespaceOption{
					Network: pb.NamespaceMode_NODE,
					// Set other fields as needed
				},
				SelinuxOptions: &pb.SELinuxOption{
					User:  "system_u",
					Role:  "system_r",
					Type:  "svirt_lxc_net_t",
					Level: "s0:c123,c456",
				},
				RunAsUser: &pb.Int64Value{
					Value: 1000, // Set to the desired user ID
				},
				RunAsGroup: &pb.Int64Value{
					Value: 1000, // Set to the desired group ID
				},
				ReadonlyRootfs: false,
				SupplementalGroups: []int64{
					1001, // Set to the desired supplemental group IDs
				},
				Privileged: true,
				// Set other fields as needed
			},
			Sysctls: map[string]string{
				"net.ipv4.ip_forward": "1",
				// Set other sysctls as needed
			},
			Overhead: &pb.LinuxContainerResources{
				CpuPeriod:          100000,
				CpuQuota:           50000,
				CpuShares:          512,
				MemoryLimitInBytes: 2147483648, // 2GB
				OomScoreAdj:        100,
				CpusetCpus:         "0-3",
				CpusetMems:         "0-1",
				HugepageLimits: []*pb.HugepageLimit{
					{
						PageSize: "2MB",
						Limit:    1073741824, // 1GB
					},
					{
						PageSize: "1GB",
						Limit:    2147483648, // 2GB
					},
				},
				Unified: map[string]string{
					"memory.max": "6937202688",
					"io.weight":  "default 100",
				},
				MemorySwapLimitInBytes: 4294967296, // 4GB
			},
			Resources: &pb.LinuxContainerResources{
				CpuPeriod:          200000,
				CpuQuota:           100000,
				CpuShares:          1024,
				MemoryLimitInBytes: 4294967296, // 4GB
				OomScoreAdj:        200,
				CpusetCpus:         "0-7",
				CpusetMems:         "0-3",
				HugepageLimits: []*pb.HugepageLimit{
					{
						PageSize: "2MB",
						Limit:    2147483648, // 2GB
					},
					{
						PageSize: "1GB",
						Limit:    4294967296, // 4GB
					},
				},
				Unified: map[string]string{
					"memory.max": "13874405376",
					"io.weight":  "default 200",
				},
				MemorySwapLimitInBytes: 8589934592, // 8GB
			},
		},
		Windows: &pb.WindowsPodSandboxConfig{
			// Set the fields of WindowsPodSandboxConfig
		},
	}

	result, err := json.Marshal(podSandboxConfig)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = os.WriteFile("podsandbox-config.json", result, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	containerConfig := pb.ContainerConfig{
		Metadata: &pb.ContainerMetadata{
			Name:    "defaultContainerName",
			Attempt: 1,
		},
		Image: &pb.ImageSpec{
			Image: "defaultImage",
		},
		Command:    []string{"defaultCommand"},
		Args:       []string{"defaultArgs"},
		WorkingDir: "defaultWorkingDir",
		Envs: []*pb.KeyValue{
			{
				Key:   "defaultEnvKey",
				Value: "defaultEnvValue",
			},
		},
		Mounts: []*pb.Mount{
			{
				ContainerPath: "defaultContainerPath",
				HostPath:      "defaultHostPath",
			},
		},
		Devices: []*pb.Device{
			{
				ContainerPath: "defaultContainerDevicePath",
				HostPath:      "defaultHostDevicePath",
			},
		},
		Labels: map[string]string{
			"defaultLabelKey": "defaultLabelValue",
		},
		Annotations: map[string]string{
			"defaultAnnotationKey": "defaultAnnotationValue",
		},
		LogPath:   "defaultLogPath",
		Stdin:     true,
		StdinOnce: true,
		Tty:       true,
		Linux: &pb.LinuxContainerConfig{
			Resources: &pb.LinuxContainerResources{
				CpuPeriod:          100000,
				CpuQuota:           50000,
				CpuShares:          512,
				MemoryLimitInBytes: 2147483648, // 2GB
				OomScoreAdj:        100,
				CpusetCpus:         "0-3",
				CpusetMems:         "0-1",
			},
			SecurityContext: &pb.LinuxContainerSecurityContext{
				Privileged: true,
				NamespaceOptions: &pb.NamespaceOption{
					Network: pb.NamespaceMode_NODE,
				},
				Capabilities: &pb.Capability{
					AddCapabilities: []string{"defaultCapability"},
				},
				ReadonlyRootfs: false,
				RunAsUser: &pb.Int64Value{
					Value: 1000, // Set to the desired user ID
				},
				RunAsGroup: &pb.Int64Value{
					Value: 1000, // Set to the desired group ID
				},
				SupplementalGroups: []int64{
					1001, // Set to the desired supplemental group IDs
				},
			},
		},
	}

	container_result, err := json.Marshal(containerConfig)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = os.WriteFile("container-config.json", container_result, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
}
