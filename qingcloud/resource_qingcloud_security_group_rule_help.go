/**
 * Copyright (c) 2016 Magicshui
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */
/**
 * Copyright (c) 2017 yunify
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package qingcloud

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	qc "github.com/yunify/qingcloud-sdk-go/service"
)

func ModifySecurityGroupRuleAttributes(d *schema.ResourceData, meta interface{}) error {
	clt := meta.(*QingCloudClient).securitygroup
	input := new(qc.ModifySecurityGroupRuleAttributesInput)
	input.SecurityGroupRule = qc.String(d.Id())
	if d.Get(resourceName).(string) != "" {
		input.SecurityGroupRuleName = qc.String(d.Get(resourceName).(string))
	} else if d.HasChange(resourceName) {
		return fmt.Errorf("name can not be modified to nil")
	} else {
		input.SecurityGroupRuleName = nil
	}
	input.Direction = qc.Int(d.Get(resourceSecurityGroupRuleDirection).(int))
	input.SecurityGroup = qc.String(d.Get(resourceSecurityGroupRuleSecurityGroupID).(string))
	input.Protocol = qc.String(d.Get(resourceSecurityGroupRuleProtocol).(string))
	input.Priority = qc.Int(d.Get(resourceSecurityGroupRulePriority).(int))
	input.RuleAction = qc.String(d.Get(resourceSecurityGroupRuleAction).(string))
	input.Val1 = getUpdateStringPointer(d, resourceSecurityGroupRuleFromPort)
	input.Val2 = getUpdateStringPointer(d, resourceSecurityGroupRuleToPort)
	input.Val3 = getUpdateStringPointer(d, resourceSecurityGroupCidrBlock)
	var output *qc.ModifySecurityGroupRuleAttributesOutput
	var err error
	simpleRetry(func() error {
		output, err = clt.ModifySecurityGroupRuleAttributes(input)
		return isServerBusy(err)
	})
	if err != nil {
		return err
	}
	return nil
}
