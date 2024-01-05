// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Popeye

package config

import (
	"fmt"
	"strconv"
	"time"
)

const (
	// DefaultUnderPerc indicates the default percentage for under allocation
	defaultUnderPerc = 200
	// DefaultOverPerc indicates the default percentage for over allocation
	defaultOverPerc = 50
)

type (
	// ID represents a sanitizer code indentifier.
	ID int

	// Glossary represents a collection of codes.
	Glossary map[ID]*Code

	// Code represents a sanitizer code.
	Code struct {
		Message  string `yaml:"message"`
		Severity Level  `yaml:"severity"`
	}

	// AllocationLimits tracks limit thresholds cpu and memory thresholds.
	AllocationLimits struct {
		CPU Allocations `yaml:"cpu"`
		MEM Allocations `yaml:"memory"`
	}

	// Allocations track under/over allocation limits.
	Allocations struct {
		UnderPerc int `yaml:"underPercUtilization"`
		OverPerc  int `yanl:"overPercUtilization"`
	}

	// Popeye tracks Popeye configuration options.
	Popeye struct {
		AllocationLimits `yaml:"allocations"`
		Excludes         `yaml:"excludes"`

		Node       Node     `yaml:"node"`
		Pod        Pod      `yaml:"pod"`
		Codes      Glossary `yaml:"codes"`
		Registries []string `yaml:"registries"`
	}

	Database struct {
		Driver      string        `mapstructure:"driver" json:"driver" yaml:"driver"`                //驱动
		Host        string        `mapstructure:"host" json:"host" yaml:"host"`                      //地址
		Port        int           `mapstructure:"port" json:"port" yaml:"port"`                      //端口
		Database    string        `mapstructure:"database" json:"database" yaml:"database"`          //数据库
		Username    string        `mapstructure:"username" json:"username" yaml:"username"`          //用户名
		Password    string        `mapstructure:"password" json:"password" yaml:"password"`          //密码
		SslMode     string        `mapstructure:"sslMode" json:"sslMode" yaml:"sslMode"`             //sslMode
		Timezone    string        `mapstructure:"timezone" json:"timezone" yaml:"timezone"`          //时区
		MaxIdle     int           `mapstructure:"maxIdle" json:"maxIdle" yaml:"maxIdle"`             //最大idle数量
		MaxOpen     int           `mapstructure:"maxOpen" json:"maxOpen" yaml:"maxOpen"`             //最大开启数量
		MaxIdleTime time.Duration `mapstructure:"maxIdleTime" json:"maxIdleTime" yaml:"maxIdleTime"` //最大idle时间
		MaxLifeTime time.Duration `mapstructure:"maxLifeTime" json:"maxLifeTime" yaml:"maxLifeTime"` //最大存活时间
	}
)

// NewPopeye create a new Popeye configuration.
func NewPopeye() Popeye {
	return Popeye{
		AllocationLimits: AllocationLimits{
			CPU: Allocations{UnderPerc: defaultUnderPerc, OverPerc: defaultOverPerc},
			MEM: Allocations{UnderPerc: defaultUnderPerc, OverPerc: defaultOverPerc},
		},
		Excludes:   newExcludes(),
		Node:       newNode(),
		Pod:        newPod(),
		Registries: []string{},
	}
}

// Format hydrates a message with arguments.
func (c *Code) Format(code ID, args ...interface{}) string {
	msg := "[POP-" + strconv.Itoa(int(code)) + "] "
	if len(args) == 0 {
		return msg + c.Message
	}
	return msg + fmt.Sprintf(c.Message, args...)
}
