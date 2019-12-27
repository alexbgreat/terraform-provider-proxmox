/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package proxmoxtf

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/danitso/terraform-provider-proxmox/proxmox"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
)

const (
	dvResourceVirtualEnvironmentVMAgentEnabled                 = false
	dvResourceVirtualEnvironmentVMAgentTrim                    = false
	dvResourceVirtualEnvironmentVMAgentType                    = "virtio"
	dvResourceVirtualEnvironmentVMCDROMEnabled                 = false
	dvResourceVirtualEnvironmentVMCDROMFileID                  = ""
	dvResourceVirtualEnvironmentVMCloudInitDNSDomain           = ""
	dvResourceVirtualEnvironmentVMCloudInitDNSServer           = ""
	dvResourceVirtualEnvironmentVMCloudInitUserAccountPassword = ""
	dvResourceVirtualEnvironmentVMCPUCores                     = 1
	dvResourceVirtualEnvironmentVMCPUHotplugged                = 0
	dvResourceVirtualEnvironmentVMCPUSockets                   = 1
	dvResourceVirtualEnvironmentVMDescription                  = ""
	dvResourceVirtualEnvironmentVMDiskDatastoreID              = "local-lvm"
	dvResourceVirtualEnvironmentVMDiskEnabled                  = true
	dvResourceVirtualEnvironmentVMDiskFileFormat               = "qcow2"
	dvResourceVirtualEnvironmentVMDiskFileID                   = ""
	dvResourceVirtualEnvironmentVMDiskSize                     = 8
	dvResourceVirtualEnvironmentVMDiskSpeedRead                = 0
	dvResourceVirtualEnvironmentVMDiskSpeedWrite               = 0
	dvResourceVirtualEnvironmentVMKeyboardLayout               = "en-us"
	dvResourceVirtualEnvironmentVMMemoryDedicated              = 512
	dvResourceVirtualEnvironmentVMMemoryFloating               = 0
	dvResourceVirtualEnvironmentVMMemoryShared                 = 0
	dvResourceVirtualEnvironmentVMName                         = ""
	dvResourceVirtualEnvironmentVMNetworkDeviceBridge          = "vmbr0"
	dvResourceVirtualEnvironmentVMNetworkDeviceEnabled         = true
	dvResourceVirtualEnvironmentVMNetworkDeviceMACAddress      = ""
	dvResourceVirtualEnvironmentVMNetworkDeviceModel           = "virtio"
	dvResourceVirtualEnvironmentVMOSType                       = "other"
	dvResourceVirtualEnvironmentVMPoolID                       = ""
	dvResourceVirtualEnvironmentVMVMID                         = -1

	mkResourceVirtualEnvironmentVMAgent                        = "agent"
	mkResourceVirtualEnvironmentVMAgentEnabled                 = "enabled"
	mkResourceVirtualEnvironmentVMAgentTrim                    = "trim"
	mkResourceVirtualEnvironmentVMAgentType                    = "type"
	mkResourceVirtualEnvironmentVMCDROM                        = "cdrom"
	mkResourceVirtualEnvironmentVMCDROMEnabled                 = "enabled"
	mkResourceVirtualEnvironmentVMCDROMFileID                  = "file_id"
	mkResourceVirtualEnvironmentVMCloudInit                    = "cloud_init"
	mkResourceVirtualEnvironmentVMCloudInitDNS                 = "dns"
	mkResourceVirtualEnvironmentVMCloudInitDNSDomain           = "domain"
	mkResourceVirtualEnvironmentVMCloudInitDNSServer           = "server"
	mkResourceVirtualEnvironmentVMCloudInitIPConfig            = "ip_config"
	mkResourceVirtualEnvironmentVMCloudInitIPConfigIPv4        = "ipv4"
	mkResourceVirtualEnvironmentVMCloudInitIPConfigIPv4Address = "address"
	mkResourceVirtualEnvironmentVMCloudInitIPConfigIPv4Gateway = "gateway"
	mkResourceVirtualEnvironmentVMCloudInitIPConfigIPv6        = "ipv6"
	mkResourceVirtualEnvironmentVMCloudInitIPConfigIPv6Address = "address"
	mkResourceVirtualEnvironmentVMCloudInitIPConfigIPv6Gateway = "gateway"
	mkResourceVirtualEnvironmentVMCloudInitUserAccount         = "user_account"
	mkResourceVirtualEnvironmentVMCloudInitUserAccountKeys     = "keys"
	mkResourceVirtualEnvironmentVMCloudInitUserAccountPassword = "password"
	mkResourceVirtualEnvironmentVMCloudInitUserAccountUsername = "username"
	mkResourceVirtualEnvironmentVMCPU                          = "cpu"
	mkResourceVirtualEnvironmentVMCPUCores                     = "cores"
	mkResourceVirtualEnvironmentVMCPUHotplugged                = "hotplugged"
	mkResourceVirtualEnvironmentVMCPUSockets                   = "sockets"
	mkResourceVirtualEnvironmentVMDescription                  = "description"
	mkResourceVirtualEnvironmentVMDisk                         = "disk"
	mkResourceVirtualEnvironmentVMDiskDatastoreID              = "datastore_id"
	mkResourceVirtualEnvironmentVMDiskEnabled                  = "enabled"
	mkResourceVirtualEnvironmentVMDiskFileFormat               = "file_format"
	mkResourceVirtualEnvironmentVMDiskFileID                   = "file_id"
	mkResourceVirtualEnvironmentVMDiskSize                     = "size"
	mkResourceVirtualEnvironmentVMDiskSpeed                    = "speed"
	mkResourceVirtualEnvironmentVMDiskSpeedRead                = "read"
	mkResourceVirtualEnvironmentVMDiskSpeedWrite               = "write"
	mkResourceVirtualEnvironmentVMKeyboardLayout               = "keyboard_layout"
	mkResourceVirtualEnvironmentVMMemory                       = "memory"
	mkResourceVirtualEnvironmentVMMemoryDedicated              = "dedicated"
	mkResourceVirtualEnvironmentVMMemoryFloating               = "floating"
	mkResourceVirtualEnvironmentVMMemoryShared                 = "shared"
	mkResourceVirtualEnvironmentVMName                         = "name"
	mkResourceVirtualEnvironmentVMNetworkDevice                = "network_device"
	mkResourceVirtualEnvironmentVMNetworkDeviceBridge          = "bridge"
	mkResourceVirtualEnvironmentVMNetworkDeviceEnabled         = "enabled"
	mkResourceVirtualEnvironmentVMNetworkDeviceMACAddress      = "mac_address"
	mkResourceVirtualEnvironmentVMNetworkDeviceModel           = "model"
	mkResourceVirtualEnvironmentVMNetworkDeviceVLANIDs         = "vlan_ids"
	mkResourceVirtualEnvironmentVMNodeName                     = "node_name"
	mkResourceVirtualEnvironmentVMOSType                       = "os_type"
	mkResourceVirtualEnvironmentVMPoolID                       = "pool_id"
	mkResourceVirtualEnvironmentVMVMID                         = "vm_id"
)

func resourceVirtualEnvironmentVM() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			mkResourceVirtualEnvironmentVMAgent: &schema.Schema{
				Type:        schema.TypeList,
				Description: "The QEMU agent configuration",
				Optional:    true,
				DefaultFunc: func() (interface{}, error) {
					defaultList := make([]interface{}, 1)
					defaultMap := make(map[string]interface{})

					defaultMap[mkResourceVirtualEnvironmentVMAgentEnabled] = dvResourceVirtualEnvironmentVMAgentEnabled
					defaultMap[mkResourceVirtualEnvironmentVMAgentTrim] = dvResourceVirtualEnvironmentVMAgentTrim
					defaultMap[mkResourceVirtualEnvironmentVMAgentType] = dvResourceVirtualEnvironmentVMAgentType

					defaultList[0] = defaultMap

					return defaultList, nil
				},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						mkResourceVirtualEnvironmentVMAgentEnabled: {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Whether to enable the QEMU agent",
							Default:     dvResourceVirtualEnvironmentVMAgentEnabled,
						},
						mkResourceVirtualEnvironmentVMAgentTrim: {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Whether to enable the FSTRIM feature in the QEMU agent",
							Default:     dvResourceVirtualEnvironmentVMAgentTrim,
						},
						mkResourceVirtualEnvironmentVMAgentType: {
							Type:         schema.TypeString,
							Optional:     true,
							Description:  "The QEMU agent interface type",
							Default:      dvResourceVirtualEnvironmentVMAgentType,
							ValidateFunc: getQEMUAgentTypeValidator(),
						},
					},
				},
				MaxItems: 1,
				MinItems: 0,
			},
			mkResourceVirtualEnvironmentVMCDROM: &schema.Schema{
				Type:        schema.TypeList,
				Description: "The CDROM drive",
				Optional:    true,
				DefaultFunc: func() (interface{}, error) {
					defaultList := make([]interface{}, 1)
					defaultMap := make(map[string]interface{})

					defaultMap[mkResourceVirtualEnvironmentVMCDROMEnabled] = dvResourceVirtualEnvironmentVMCDROMEnabled
					defaultMap[mkResourceVirtualEnvironmentVMCDROMFileID] = dvResourceVirtualEnvironmentVMCDROMFileID

					defaultList[0] = defaultMap

					return defaultList, nil
				},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						mkResourceVirtualEnvironmentVMCDROMEnabled: {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Whether to enable the CDROM drive",
							Default:     dvResourceVirtualEnvironmentVMCDROMEnabled,
						},
						mkResourceVirtualEnvironmentVMCDROMFileID: {
							Type:         schema.TypeString,
							Optional:     true,
							Description:  "The file id",
							Default:      dvResourceVirtualEnvironmentVMCDROMFileID,
							ValidateFunc: getFileIDValidator(),
						},
					},
				},
				MaxItems: 1,
				MinItems: 0,
			},
			mkResourceVirtualEnvironmentVMCloudInit: &schema.Schema{
				Type:        schema.TypeList,
				Description: "The cloud-init configuration",
				Optional:    true,
				DefaultFunc: func() (interface{}, error) {
					return make([]interface{}, 0), nil
				},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						mkResourceVirtualEnvironmentVMCloudInitDNS: {
							Type:        schema.TypeList,
							Description: "The DNS configuration",
							Optional:    true,
							DefaultFunc: func() (interface{}, error) {
								return make([]interface{}, 0), nil
							},
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									mkResourceVirtualEnvironmentVMCloudInitDNSDomain: {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The DNS search domain",
										Default:     dvResourceVirtualEnvironmentVMCloudInitDNSDomain,
									},
									mkResourceVirtualEnvironmentVMCloudInitDNSServer: {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The DNS server",
										Default:     dvResourceVirtualEnvironmentVMCloudInitDNSServer,
									},
								},
							},
							MaxItems: 1,
							MinItems: 0,
						},
						mkResourceVirtualEnvironmentVMCloudInitIPConfig: {
							Type:        schema.TypeList,
							Description: "The IP configuration",
							Optional:    true,
							DefaultFunc: func() (interface{}, error) {
								return make([]interface{}, 0), nil
							},
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									mkResourceVirtualEnvironmentVMCloudInitIPConfigIPv4: {
										Type:        schema.TypeList,
										Description: "The IPv4 configuration",
										Optional:    true,
										DefaultFunc: func() (interface{}, error) {
											return make([]interface{}, 0), nil
										},
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												mkResourceVirtualEnvironmentVMCloudInitIPConfigIPv4Address: {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "The IPv4 address",
													Default:     "",
												},
												mkResourceVirtualEnvironmentVMCloudInitIPConfigIPv4Gateway: {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "The IPv4 gateway",
													Default:     "",
												},
											},
										},
										MaxItems: 1,
										MinItems: 0,
									},
									mkResourceVirtualEnvironmentVMCloudInitIPConfigIPv6: {
										Type:        schema.TypeList,
										Description: "The IPv6 configuration",
										Optional:    true,
										DefaultFunc: func() (interface{}, error) {
											return make([]interface{}, 0), nil
										},
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												mkResourceVirtualEnvironmentVMCloudInitIPConfigIPv6Address: {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "The IPv6 address",
													Default:     "",
												},
												mkResourceVirtualEnvironmentVMCloudInitIPConfigIPv6Gateway: {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "The IPv6 gateway",
													Default:     "",
												},
											},
										},
										MaxItems: 1,
										MinItems: 0,
									},
								},
							},
							MaxItems: 8,
							MinItems: 0,
						},
						mkResourceVirtualEnvironmentVMCloudInitUserAccount: {
							Type:        schema.TypeList,
							Description: "The user account configuration",
							Required:    true,
							DefaultFunc: func() (interface{}, error) {
								return make([]interface{}, 0), nil
							},
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									mkResourceVirtualEnvironmentVMCloudInitUserAccountKeys: {
										Type:        schema.TypeList,
										Required:    true,
										Description: "The SSH keys",
										Elem:        &schema.Schema{Type: schema.TypeString},
									},
									mkResourceVirtualEnvironmentVMCloudInitUserAccountPassword: {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The SSH password",
										Default:     dvResourceVirtualEnvironmentVMCloudInitUserAccountPassword,
									},
									mkResourceVirtualEnvironmentVMCloudInitUserAccountUsername: {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The SSH username",
									},
								},
							},
							MaxItems: 1,
							MinItems: 0,
						},
					},
				},
				MaxItems: 1,
				MinItems: 0,
			},
			mkResourceVirtualEnvironmentVMCPU: &schema.Schema{
				Type:        schema.TypeList,
				Description: "The CPU allocation",
				Optional:    true,
				DefaultFunc: func() (interface{}, error) {
					defaultList := make([]interface{}, 1)
					defaultMap := make(map[string]interface{})

					defaultMap[mkResourceVirtualEnvironmentVMCPUCores] = dvResourceVirtualEnvironmentVMCPUCores
					defaultMap[mkResourceVirtualEnvironmentVMCPUHotplugged] = dvResourceVirtualEnvironmentVMCPUHotplugged
					defaultMap[mkResourceVirtualEnvironmentVMCPUSockets] = dvResourceVirtualEnvironmentVMCPUSockets

					defaultList[0] = defaultMap

					return defaultList, nil
				},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						mkResourceVirtualEnvironmentVMCPUCores: {
							Type:         schema.TypeInt,
							Optional:     true,
							Description:  "The number of CPU cores",
							Default:      dvResourceVirtualEnvironmentVMCPUCores,
							ValidateFunc: validation.IntBetween(1, 2304),
						},
						mkResourceVirtualEnvironmentVMCPUHotplugged: {
							Type:         schema.TypeInt,
							Optional:     true,
							Description:  "The number of hotplugged vCPUs",
							Default:      dvResourceVirtualEnvironmentVMCPUHotplugged,
							ValidateFunc: validation.IntBetween(0, 2304),
						},
						mkResourceVirtualEnvironmentVMCPUSockets: {
							Type:         schema.TypeInt,
							Optional:     true,
							Description:  "The number of CPU sockets",
							Default:      dvResourceVirtualEnvironmentVMCPUSockets,
							ValidateFunc: validation.IntBetween(1, 16),
						},
					},
				},
				MaxItems: 1,
				MinItems: 0,
			},
			mkResourceVirtualEnvironmentVMDescription: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description",
				Default:     dvResourceVirtualEnvironmentVMDescription,
			},
			mkResourceVirtualEnvironmentVMDisk: &schema.Schema{
				Type:        schema.TypeList,
				Description: "The disk devices",
				Optional:    true,
				DefaultFunc: func() (interface{}, error) {
					defaultList := make([]interface{}, 1)
					defaultMap := make(map[string]interface{})

					defaultMap[mkResourceVirtualEnvironmentVMDiskDatastoreID] = dvResourceVirtualEnvironmentVMDiskDatastoreID
					defaultMap[mkResourceVirtualEnvironmentVMDiskFileFormat] = dvResourceVirtualEnvironmentVMDiskFileFormat
					defaultMap[mkResourceVirtualEnvironmentVMDiskFileID] = dvResourceVirtualEnvironmentVMDiskFileID
					defaultMap[mkResourceVirtualEnvironmentVMDiskSize] = dvResourceVirtualEnvironmentVMDiskSize

					defaultList[0] = defaultMap

					return defaultList, nil
				},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						mkResourceVirtualEnvironmentVMDiskDatastoreID: {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The datastore id",
							Default:     dvResourceVirtualEnvironmentVMDiskDatastoreID,
						},
						mkResourceVirtualEnvironmentVMDiskEnabled: {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Whether to enable the disk",
							Default:     dvResourceVirtualEnvironmentVMDiskEnabled,
						},
						mkResourceVirtualEnvironmentVMDiskFileFormat: {
							Type:         schema.TypeString,
							Optional:     true,
							Description:  "The file format",
							Default:      dvResourceVirtualEnvironmentVMDiskFileFormat,
							ValidateFunc: getFileFormatValidator(),
						},
						mkResourceVirtualEnvironmentVMDiskFileID: {
							Type:         schema.TypeString,
							Optional:     true,
							Description:  "The file id for a disk image",
							Default:      dvResourceVirtualEnvironmentVMDiskFileID,
							ValidateFunc: getFileIDValidator(),
						},
						mkResourceVirtualEnvironmentVMDiskSize: {
							Type:         schema.TypeInt,
							Optional:     true,
							Description:  "The disk size in gigabytes",
							Default:      dvResourceVirtualEnvironmentVMDiskSize,
							ValidateFunc: validation.IntBetween(1, 8192),
						},
						mkResourceVirtualEnvironmentVMDiskSpeed: {
							Type:        schema.TypeList,
							Description: "The speed limits",
							Optional:    true,
							DefaultFunc: func() (interface{}, error) {
								defaultList := make([]interface{}, 1)
								defaultMap := make(map[string]interface{})

								defaultMap[mkResourceVirtualEnvironmentVMDiskSpeedRead] = dvResourceVirtualEnvironmentVMDiskSpeedRead
								defaultMap[mkResourceVirtualEnvironmentVMDiskSpeedWrite] = dvResourceVirtualEnvironmentVMDiskSpeedWrite

								defaultList[0] = defaultMap

								return defaultList, nil
							},
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									mkResourceVirtualEnvironmentVMDiskSpeedRead: {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "The maximum read speed in megabytes per second",
										Default:     dvResourceVirtualEnvironmentVMDiskSpeedRead,
									},
									mkResourceVirtualEnvironmentVMDiskSpeedWrite: {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "The maximum write speed in megabytes per second",
										Default:     dvResourceVirtualEnvironmentVMDiskSpeedRead,
									},
								},
							},
							MaxItems: 1,
							MinItems: 0,
						},
					},
				},
				MaxItems: 14,
				MinItems: 0,
			},
			mkResourceVirtualEnvironmentVMKeyboardLayout: {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  "The keyboard layout",
				Default:      dvResourceVirtualEnvironmentVMKeyboardLayout,
				ValidateFunc: getKeyboardLayoutValidator(),
			},
			mkResourceVirtualEnvironmentVMMemory: &schema.Schema{
				Type:        schema.TypeList,
				Description: "The memory allocation",
				Optional:    true,
				DefaultFunc: func() (interface{}, error) {
					defaultList := make([]interface{}, 1)
					defaultMap := make(map[string]interface{})

					defaultMap[mkResourceVirtualEnvironmentVMMemoryDedicated] = dvResourceVirtualEnvironmentVMMemoryDedicated
					defaultMap[mkResourceVirtualEnvironmentVMMemoryFloating] = dvResourceVirtualEnvironmentVMMemoryFloating
					defaultMap[mkResourceVirtualEnvironmentVMMemoryShared] = dvResourceVirtualEnvironmentVMMemoryShared

					defaultList[0] = defaultMap

					return defaultList, nil
				},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						mkResourceVirtualEnvironmentVMMemoryDedicated: {
							Type:         schema.TypeInt,
							Optional:     true,
							Description:  "The dedicated memory in megabytes",
							Default:      dvResourceVirtualEnvironmentVMMemoryDedicated,
							ValidateFunc: validation.IntBetween(64, 268435456),
						},
						mkResourceVirtualEnvironmentVMMemoryFloating: {
							Type:         schema.TypeInt,
							Optional:     true,
							Description:  "The floating memory in megabytes (balloon)",
							Default:      dvResourceVirtualEnvironmentVMMemoryFloating,
							ValidateFunc: validation.IntBetween(0, 268435456),
						},
						mkResourceVirtualEnvironmentVMMemoryShared: {
							Type:         schema.TypeInt,
							Optional:     true,
							Description:  "The shared memory in megabytes",
							Default:      dvResourceVirtualEnvironmentVMMemoryShared,
							ValidateFunc: validation.IntBetween(0, 268435456),
						},
					},
				},
				MaxItems: 1,
				MinItems: 0,
			},
			mkResourceVirtualEnvironmentVMName: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name",
				Default:     dvResourceVirtualEnvironmentVMName,
			},
			mkResourceVirtualEnvironmentVMNetworkDevice: &schema.Schema{
				Type:        schema.TypeList,
				Description: "The network devices",
				Optional:    true,
				DefaultFunc: func() (interface{}, error) {
					return make([]interface{}, 1), nil
				},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						mkResourceVirtualEnvironmentVMNetworkDeviceBridge: {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The bridge",
							Default:     dvResourceVirtualEnvironmentVMNetworkDeviceBridge,
						},
						mkResourceVirtualEnvironmentVMNetworkDeviceEnabled: {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Whether to enable the network device",
							Default:     dvResourceVirtualEnvironmentVMNetworkDeviceEnabled,
						},
						mkResourceVirtualEnvironmentVMNetworkDeviceMACAddress: {
							Type:         schema.TypeString,
							Optional:     true,
							Description:  "The MAC address",
							Default:      dvResourceVirtualEnvironmentVMNetworkDeviceMACAddress,
							ValidateFunc: getMACAddressValidator(),
						},
						mkResourceVirtualEnvironmentVMNetworkDeviceModel: {
							Type:         schema.TypeString,
							Optional:     true,
							Description:  "The model",
							Default:      dvResourceVirtualEnvironmentVMNetworkDeviceModel,
							ValidateFunc: getNetworkDeviceModelValidator(),
						},
						mkResourceVirtualEnvironmentVMNetworkDeviceVLANIDs: {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "The VLAN identifiers",
							DefaultFunc: func() (interface{}, error) {
								return make([]interface{}, 0), nil
							},
							Elem: &schema.Schema{Type: schema.TypeInt},
						},
					},
				},
				MaxItems: 8,
				MinItems: 0,
			},
			mkResourceVirtualEnvironmentVMNodeName: &schema.Schema{
				Type:        schema.TypeString,
				Description: "The node name",
				Required:    true,
			},
			mkResourceVirtualEnvironmentVMOSType: {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  "The OS type",
				Default:      dvResourceVirtualEnvironmentVMOSType,
				ValidateFunc: getOSTypeValidator(),
			},
			mkResourceVirtualEnvironmentVMPoolID: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The ID of the pool to assign the virtual machine to",
				Default:     dvResourceVirtualEnvironmentVMPoolID,
			},
			mkResourceVirtualEnvironmentVMVMID: {
				Type:         schema.TypeInt,
				Optional:     true,
				ForceNew:     true,
				Description:  "The VM identifier",
				Default:      dvResourceVirtualEnvironmentVMVMID,
				ValidateFunc: getVMIDValidator(),
			},
		},
		Create: resourceVirtualEnvironmentVMCreate,
		Read:   resourceVirtualEnvironmentVMRead,
		Update: resourceVirtualEnvironmentVMUpdate,
		Delete: resourceVirtualEnvironmentVMDelete,
	}
}

func resourceVirtualEnvironmentVMCreate(d *schema.ResourceData, m interface{}) error {
	config := m.(providerConfiguration)
	veClient, err := config.GetVEClient()

	if err != nil {
		return err
	}

	resourceSchema := resourceVirtualEnvironmentVM().Schema
	agent := d.Get(mkResourceVirtualEnvironmentVMAgent).([]interface{})

	if len(agent) == 0 {
		agentDefault, err := resourceSchema[mkResourceVirtualEnvironmentVMAgent].DefaultValue()

		if err != nil {
			return err
		}

		agent = agentDefault.([]interface{})
	}

	agentBlock := agent[0].(map[string]interface{})
	agentEnabled := proxmox.CustomBool(agentBlock[mkResourceVirtualEnvironmentVMAgentEnabled].(bool))
	agentTrim := proxmox.CustomBool(agentBlock[mkResourceVirtualEnvironmentVMAgentTrim].(bool))
	agentType := agentBlock[mkResourceVirtualEnvironmentVMAgentType].(string)

	cdrom := d.Get(mkResourceVirtualEnvironmentVMCDROM).([]interface{})

	if len(cdrom) == 0 {
		cdromDefault, err := resourceSchema[mkResourceVirtualEnvironmentVMCDROM].DefaultValue()

		if err != nil {
			return err
		}

		cdrom = cdromDefault.([]interface{})
	}

	cdromBlock := cdrom[0].(map[string]interface{})
	cdromEnabled := cdromBlock[mkResourceVirtualEnvironmentVMCDROMEnabled].(bool)
	cdromFileID := cdromBlock[mkResourceVirtualEnvironmentVMCDROMFileID].(string)

	if cdromFileID == "" {
		cdromFileID = "cdrom"
	}

	var cloudInitConfig *proxmox.CustomCloudInitConfig

	cloudInit := d.Get(mkResourceVirtualEnvironmentVMCloudInit).([]interface{})

	if len(cloudInit) > 0 {
		cdromEnabled = true
		cdromFileID = "local-lvm:cloudinit"

		cloudInitBlock := cloudInit[0].(map[string]interface{})
		cloudInitConfig = &proxmox.CustomCloudInitConfig{}
		cloudInitDNS := cloudInitBlock[mkResourceVirtualEnvironmentVMCloudInitDNS].([]interface{})

		if len(cloudInitDNS) > 0 {
			cloudInitDNSBlock := cloudInitDNS[0].(map[string]interface{})
			domain := cloudInitDNSBlock[mkResourceVirtualEnvironmentVMCloudInitDNSDomain].(string)

			if domain != "" {
				cloudInitConfig.SearchDomain = &domain
			}

			server := cloudInitDNSBlock[mkResourceVirtualEnvironmentVMCloudInitDNSServer].(string)

			if server != "" {
				cloudInitConfig.Nameserver = &server
			}
		}

		cloudInitIPConfig := cloudInitBlock[mkResourceVirtualEnvironmentVMCloudInitIPConfig].([]interface{})
		cloudInitConfig.IPConfig = make([]proxmox.CustomCloudInitIPConfig, len(cloudInitIPConfig))

		for i, c := range cloudInitIPConfig {
			configBlock := c.(map[string]interface{})
			ipv4 := configBlock[mkResourceVirtualEnvironmentVMCloudInitIPConfigIPv4].([]interface{})

			if len(ipv4) > 0 {
				ipv4Block := ipv4[0].(map[string]interface{})
				ipv4Address := ipv4Block[mkResourceVirtualEnvironmentVMCloudInitIPConfigIPv4Address].(string)

				if ipv4Address != "" {
					cloudInitConfig.IPConfig[i].IPv4 = &ipv4Address
				}

				ipv4Gateway := ipv4Block[mkResourceVirtualEnvironmentVMCloudInitIPConfigIPv4Gateway].(string)

				if ipv4Gateway != "" {
					cloudInitConfig.IPConfig[i].GatewayIPv4 = &ipv4Gateway
				}
			}

			ipv6 := configBlock[mkResourceVirtualEnvironmentVMCloudInitIPConfigIPv6].([]interface{})

			if len(ipv6) > 0 {
				ipv6Block := ipv6[0].(map[string]interface{})
				ipv6Address := ipv6Block[mkResourceVirtualEnvironmentVMCloudInitIPConfigIPv6Address].(string)

				if ipv6Address != "" {
					cloudInitConfig.IPConfig[i].IPv6 = &ipv6Address
				}

				ipv6Gateway := ipv6Block[mkResourceVirtualEnvironmentVMCloudInitIPConfigIPv6Gateway].(string)

				if ipv6Gateway != "" {
					cloudInitConfig.IPConfig[i].GatewayIPv6 = &ipv6Gateway
				}
			}
		}

		cloudInitUserAccount := cloudInitBlock[mkResourceVirtualEnvironmentVMCloudInitUserAccount].([]interface{})

		if len(cloudInitUserAccount) > 0 {
			cloudInitUserAccountBlock := cloudInitUserAccount[0].(map[string]interface{})
			keys := cloudInitUserAccountBlock[mkResourceVirtualEnvironmentVMCloudInitUserAccountKeys].([]interface{})

			if len(keys) > 0 {
				sshKeys := make(proxmox.CustomCloudInitSSHKeys, len(keys))

				for i, k := range keys {
					sshKeys[i] = k.(string)
				}

				cloudInitConfig.SSHKeys = &sshKeys
			}

			password := cloudInitUserAccountBlock[mkResourceVirtualEnvironmentVMCloudInitUserAccountPassword].(string)

			if password != "" {
				cloudInitConfig.Password = &password
			}

			username := cloudInitUserAccountBlock[mkResourceVirtualEnvironmentVMCloudInitUserAccountUsername].(string)

			cloudInitConfig.Username = &username
		}
	}

	cpu := d.Get(mkResourceVirtualEnvironmentVMCPU).([]interface{})

	if len(cpu) == 0 {
		cpuDefault, err := resourceSchema[mkResourceVirtualEnvironmentVMCPU].DefaultValue()

		if err != nil {
			return err
		}

		cpu = cpuDefault.([]interface{})
	}

	cpuBlock := cpu[0].(map[string]interface{})
	cpuCores := cpuBlock[mkResourceVirtualEnvironmentVMCPUCores].(int)
	cpuHotplugged := cpuBlock[mkResourceVirtualEnvironmentVMCPUHotplugged].(int)
	cpuSockets := cpuBlock[mkResourceVirtualEnvironmentVMCPUSockets].(int)

	description := d.Get(mkResourceVirtualEnvironmentVMDescription).(string)
	disk := d.Get(mkResourceVirtualEnvironmentVMDisk).([]interface{})
	scsiDevices := make(proxmox.CustomStorageDevices, len(disk))

	diskSchemaElem := resourceSchema[mkResourceVirtualEnvironmentVMDisk].Elem
	diskSchemaResource := diskSchemaElem.(*schema.Resource)
	diskSpeedResource := diskSchemaResource.Schema[mkResourceVirtualEnvironmentVMDiskSpeed]

	for i, d := range disk {
		block := d.(map[string]interface{})

		datastoreID, _ := block[mkResourceVirtualEnvironmentVMDiskDatastoreID].(string)
		enabled, _ := block[mkResourceVirtualEnvironmentVMDiskEnabled].(bool)
		fileID, _ := block[mkResourceVirtualEnvironmentVMDiskFileID].(string)
		size, _ := block[mkResourceVirtualEnvironmentVMDiskSize].(int)
		speed := block[mkResourceVirtualEnvironmentVMDiskSpeed].([]interface{})

		if len(speed) == 0 {
			diskSpeedDefault, err := diskSpeedResource.DefaultValue()

			if err != nil {
				return err
			}

			speed = diskSpeedDefault.([]interface{})
		}

		speedBlock := speed[0].(map[string]interface{})
		speedLimitRead := speedBlock[mkResourceVirtualEnvironmentVMDiskSpeedRead].(int)
		speedLimitWrite := speedBlock[mkResourceVirtualEnvironmentVMDiskSpeedWrite].(int)

		diskDevice := proxmox.CustomStorageDevice{
			Enabled: enabled,
		}

		if speedLimitRead > 0 {
			diskDevice.MaxReadSpeedMbps = &speedLimitRead
		}

		if speedLimitWrite > 0 {
			diskDevice.MaxWriteSpeedMbps = &speedLimitWrite
		}

		if fileID != "" {
			diskDevice.Enabled = false
		} else {
			diskDevice.FileVolume = fmt.Sprintf("%s:%d", datastoreID, size)
		}

		scsiDevices[i] = diskDevice
	}

	keyboardLayout := d.Get(mkResourceVirtualEnvironmentVMKeyboardLayout).(string)
	memory := d.Get(mkResourceVirtualEnvironmentVMMemory).([]interface{})

	if len(memory) == 0 {
		memoryDefault, err := resourceSchema[mkResourceVirtualEnvironmentVMMemory].DefaultValue()

		if err != nil {
			return err
		}

		memory = memoryDefault.([]interface{})
	}

	memoryBlock := memory[0].(map[string]interface{})
	memoryDedicated := memoryBlock[mkResourceVirtualEnvironmentVMMemoryDedicated].(int)
	memoryFloating := memoryBlock[mkResourceVirtualEnvironmentVMMemoryFloating].(int)
	memoryShared := memoryBlock[mkResourceVirtualEnvironmentVMMemoryShared].(int)

	name := d.Get(mkResourceVirtualEnvironmentVMName).(string)

	networkDevice := d.Get(mkResourceVirtualEnvironmentVMNetworkDevice).([]interface{})
	networkDeviceObjects := make(proxmox.CustomNetworkDevices, len(networkDevice))

	for i, d := range networkDevice {
		block := d.(map[string]interface{})

		bridge, _ := block[mkResourceVirtualEnvironmentVMNetworkDeviceBridge].(string)
		enabled, _ := block[mkResourceVirtualEnvironmentVMNetworkDeviceEnabled].(bool)
		macAddress, _ := block[mkResourceVirtualEnvironmentVMNetworkDeviceMACAddress].(string)
		model, _ := block[mkResourceVirtualEnvironmentVMNetworkDeviceModel].(string)
		vlanIDs, _ := block[mkResourceVirtualEnvironmentVMNetworkDeviceVLANIDs].([]interface{})

		device := proxmox.CustomNetworkDevice{
			Enabled: enabled,
			Model:   model,
		}

		if bridge != "" {
			device.Bridge = &bridge
		}

		if macAddress != "" {
			device.MACAddress = &macAddress
		}

		if len(vlanIDs) > 0 {
			device.Trunks = make([]int, len(vlanIDs))

			for vi, vv := range vlanIDs {
				device.Trunks[vi] = vv.(int)
			}
		}

		networkDeviceObjects[i] = device
	}

	nodeName := d.Get(mkResourceVirtualEnvironmentVMNodeName).(string)
	osType := d.Get(mkResourceVirtualEnvironmentVMOSType).(string)
	poolID := d.Get(mkResourceVirtualEnvironmentVMPoolID).(string)
	vmID := d.Get(mkResourceVirtualEnvironmentVMVMID).(int)

	if vmID == -1 {
		vmIDNew, err := veClient.GetVMID()

		if err != nil {
			return err
		}

		vmID = *vmIDNew
	}

	var memorySharedObject *proxmox.CustomSharedMemory

	bootDisk := "scsi0"
	bootOrder := "c"

	if cdromEnabled {
		bootOrder = "cd"
	}

	ideDevice2Media := "cdrom"
	ideDevices := proxmox.CustomStorageDevices{
		proxmox.CustomStorageDevice{
			Enabled: false,
		},
		proxmox.CustomStorageDevice{
			Enabled: false,
		},
		proxmox.CustomStorageDevice{
			Enabled:    cdromEnabled,
			FileVolume: cdromFileID,
			Media:      &ideDevice2Media,
		},
	}

	if memoryShared > 0 {
		memorySharedName := fmt.Sprintf("vm-%d-ivshmem", vmID)
		memorySharedObject = &proxmox.CustomSharedMemory{
			Name: &memorySharedName,
			Size: memoryShared,
		}
	}

	scsiHardware := "virtio-scsi-pci"
	startOnBoot := proxmox.CustomBool(true)
	tabletDeviceEnabled := proxmox.CustomBool(true)

	body := &proxmox.VirtualEnvironmentVMCreateRequestBody{
		Agent: &proxmox.CustomAgent{
			Enabled:         &agentEnabled,
			TrimClonedDisks: &agentTrim,
			Type:            &agentType,
		},
		BootDisk:            &bootDisk,
		BootOrder:           &bootOrder,
		CloudInitConfig:     cloudInitConfig,
		CPUCores:            &cpuCores,
		CPUSockets:          &cpuSockets,
		DedicatedMemory:     &memoryDedicated,
		FloatingMemory:      &memoryFloating,
		IDEDevices:          ideDevices,
		KeyboardLayout:      &keyboardLayout,
		NetworkDevices:      networkDeviceObjects,
		OSType:              &osType,
		SCSIDevices:         scsiDevices,
		SCSIHardware:        &scsiHardware,
		SerialDevices:       []string{"socket"},
		SharedMemory:        memorySharedObject,
		StartOnBoot:         &startOnBoot,
		TabletDeviceEnabled: &tabletDeviceEnabled,
		VMID:                &vmID,
	}

	if cpuHotplugged > 0 {
		body.VirtualCPUCount = &cpuHotplugged
	}

	if description != "" {
		body.Description = &description
	}

	if name != "" {
		body.Name = &name
	}

	if poolID != "" {
		body.PoolID = &poolID
	}

	err = veClient.CreateVM(nodeName, body)

	if err != nil {
		return err
	}

	d.SetId(strconv.Itoa(vmID))

	return resourceVirtualEnvironmentVMCreateImportedDisks(d, m)
}

func resourceVirtualEnvironmentVMCreateImportedDisks(d *schema.ResourceData, m interface{}) error {
	config := m.(providerConfiguration)
	veClient, err := config.GetVEClient()

	if err != nil {
		return err
	}

	nodeName := d.Get(mkResourceVirtualEnvironmentVMNodeName).(string)
	vmID, err := strconv.Atoi(d.Id())

	if err != nil {
		return err
	}

	commands := []string{}

	// Determine the ID of the next disk.
	disk := d.Get(mkResourceVirtualEnvironmentVMDisk).([]interface{})
	diskCount := 0

	for _, d := range disk {
		block := d.(map[string]interface{})
		fileID, _ := block[mkResourceVirtualEnvironmentVMDiskFileID].(string)

		if fileID == "" {
			diskCount++
		}
	}

	// Retrieve some information about the disk schema.
	resourceSchema := resourceVirtualEnvironmentVM().Schema
	diskSchemaElem := resourceSchema[mkResourceVirtualEnvironmentVMDisk].Elem
	diskSchemaResource := diskSchemaElem.(*schema.Resource)
	diskSpeedResource := diskSchemaResource.Schema[mkResourceVirtualEnvironmentVMDiskSpeed]

	// Generate the commands required to import the specified disks.
	importedDiskCount := 0

	for i, d := range disk {
		block := d.(map[string]interface{})

		enabled, _ := block[mkResourceVirtualEnvironmentVMDiskEnabled].(bool)
		fileID, _ := block[mkResourceVirtualEnvironmentVMDiskFileID].(string)

		if !enabled || fileID == "" {
			continue
		}

		datastoreID, _ := block[mkResourceVirtualEnvironmentVMDiskDatastoreID].(string)
		fileFormat, _ := block[mkResourceVirtualEnvironmentVMDiskFileFormat].(string)
		size, _ := block[mkResourceVirtualEnvironmentVMDiskSize].(int)
		speed := block[mkResourceVirtualEnvironmentVMDiskSpeed].([]interface{})

		if len(speed) == 0 {
			diskSpeedDefault, err := diskSpeedResource.DefaultValue()

			if err != nil {
				return err
			}

			speed = diskSpeedDefault.([]interface{})
		}

		speedBlock := speed[0].(map[string]interface{})
		speedLimitRead := speedBlock[mkResourceVirtualEnvironmentVMDiskSpeedRead].(int)
		speedLimitWrite := speedBlock[mkResourceVirtualEnvironmentVMDiskSpeedWrite].(int)

		diskOptions := ""

		if speedLimitRead > 0 {
			diskOptions += fmt.Sprintf(",mbps_rd=%d", speedLimitRead)
		}

		if speedLimitWrite > 0 {
			diskOptions += fmt.Sprintf(",mbps_wr=%d", speedLimitWrite)
		}

		fileIDParts := strings.Split(fileID, ":")
		filePath := fmt.Sprintf("/var/lib/vz/template/%s", fileIDParts[1])
		filePathTmp := fmt.Sprintf("/tmp/vm-%d-disk-%d.%s", vmID, diskCount+importedDiskCount, fileFormat)

		commands = append(
			commands,
			fmt.Sprintf("cp %s %s", filePath, filePathTmp),
			fmt.Sprintf("qemu-img resize %s %dG", filePathTmp, size),
			fmt.Sprintf("qm importdisk %d %s %s -format qcow2", vmID, filePathTmp, datastoreID),
			fmt.Sprintf("qm set %d -scsi%d %s:vm-%d-disk-%d%s", vmID, i, datastoreID, vmID, diskCount+importedDiskCount, diskOptions),
			fmt.Sprintf("rm -f %s", filePathTmp),
		)

		importedDiskCount++
	}

	// Execute the commands on the node and wait for the result.
	// This is a highly experimental approach to disk imports and is not recommended by Proxmox.
	if len(commands) > 0 {
		err = veClient.ExecuteNodeCommands(nodeName, commands)

		if err != nil {
			return err
		}
	}

	return resourceVirtualEnvironmentVMCreateStart(d, m)
}

func resourceVirtualEnvironmentVMCreateStart(d *schema.ResourceData, m interface{}) error {
	config := m.(providerConfiguration)
	veClient, err := config.GetVEClient()

	if err != nil {
		return err
	}

	nodeName := d.Get(mkResourceVirtualEnvironmentVMNodeName).(string)
	vmID, err := strconv.Atoi(d.Id())

	if err != nil {
		return err
	}

	// Start the virtual machine and wait for it to reach a running state before continuing.
	err = veClient.StartVM(nodeName, vmID)

	if err != nil {
		return err
	}

	err = veClient.WaitForState(nodeName, vmID, "running", 120, 5)

	if err != nil {
		return err
	}

	return resourceVirtualEnvironmentVMRead(d, m)
}

func resourceVirtualEnvironmentVMRead(d *schema.ResourceData, m interface{}) error {
	/*
		config := m.(providerConfiguration)
		veClient, err := config.GetVEClient()

		if err != nil {
			return err
		}
	*/

	return nil
}

func resourceVirtualEnvironmentVMUpdate(d *schema.ResourceData, m interface{}) error {
	/*
		config := m.(providerConfiguration)
		veClient, err := config.GetVEClient()

		if err != nil {
			return err
		}
	*/

	return resourceVirtualEnvironmentVMRead(d, m)
}

func resourceVirtualEnvironmentVMDelete(d *schema.ResourceData, m interface{}) error {
	config := m.(providerConfiguration)
	veClient, err := config.GetVEClient()

	if err != nil {
		return err
	}

	nodeName := d.Get(mkResourceVirtualEnvironmentVMNodeName).(string)
	vmID, err := strconv.Atoi(d.Id())

	if err != nil {
		return err
	}

	// Shut down the virtual machine before deleting it.
	forceStop := proxmox.CustomBool(true)
	shutdownTimeout := 300

	err = veClient.ShutdownVM(nodeName, vmID, &proxmox.VirtualEnvironmentVMShutdownRequestBody{
		ForceStop: &forceStop,
		Timeout:   &shutdownTimeout,
	})

	if err != nil {
		return err
	}

	err = veClient.WaitForState(nodeName, vmID, "stopped", 30, 5)

	if err != nil {
		return err
	}

	err = veClient.DeleteVM(nodeName, vmID)

	if err != nil {
		if strings.Contains(err.Error(), "HTTP 404") {
			d.SetId("")

			return nil
		}

		return err
	}

	// Wait for the state to become unavailable as that clearly indicates the destruction of the VM.
	err = veClient.WaitForState(nodeName, vmID, "", 30, 2)

	if err == nil {
		return fmt.Errorf("Failed to delete VM \"%d\"", vmID)
	}

	d.SetId("")

	return nil
}