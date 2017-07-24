package main

import (
	"log"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/quintilesims/layer0/cli/client"
)

func resourceLayer0Service() *schema.Resource {
	return &schema.Resource{
		Create: resourceLayer0ServiceCreate,
		Read:   resourceLayer0ServiceRead,
		Update: resourceLayer0ServiceUpdate,
		Delete: resourceLayer0ServiceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"environment": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"deploy": {
				Type:     schema.TypeString,
				Required: true,
			},
			"load_balancer": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"scale": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
		},
	}
}

func resourceLayer0ServiceCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(client.Client)

	environmentID := d.Get("environment").(string)
	name := d.Get("name").(string)
	deployID := d.Get("deploy").(string)
	loadBalancerID := d.Get("load_balancer").(string)
	scale := d.Get("scale").(int)

	service, err := client.CreateService(name, environmentID, deployID, loadBalancerID)
	if err != nil {
		return err
	}

	// set id first to tell terraform resource has been created
	d.SetId(service.ServiceID)

	if scale != 1 {
		if _, err := client.ScaleService(service.ServiceID, scale); err != nil {
			return err
		}
	}

	if _, err := client.WaitForDeployment(service.ServiceID, defaultTimeout); err != nil {
		return err
	}

	return resourceLayer0ServiceRead(d, meta)
}

func resourceLayer0ServiceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(client.Client)
	serviceID := d.Id()

	service, err := client.GetService(serviceID)
	if err != nil {
		if strings.Contains(err.Error(), "No service found") {
			d.SetId("")
			log.Printf("[WARN] Error Reading Service (%s), service does not exist", serviceID)
			return nil
		}

		return err
	}

	d.Set("environment", service.EnvironmentID)
	d.Set("name", service.ServiceName)
	d.Set("load_balancer", service.LoadBalancerID)
	d.Set("scale", service.DesiredCount)

	for _, deployment := range service.Deployments {
		if deployment.Status == "PRIMARY" {
			d.Set("deploy", deployment.DeployID)
		}
	}

	return nil
}

func resourceLayer0ServiceUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(client.Client)
	serviceID := d.Id()

	d.Partial(true)

	if d.HasChange("deploy") {
		deployID := d.Get("deploy").(string)

		if _, err := client.UpdateService(serviceID, deployID); err != nil {
			return err
		}

		d.SetPartial("deploy")
	}

	if d.HasChange("scale") {
		scale := d.Get("scale").(int)

		if _, err := client.ScaleService(serviceID, scale); err != nil {
			return err
		}

		d.SetPartial("scale")
	}

	d.Partial(false)

	if _, err := client.WaitForDeployment(serviceID, defaultTimeout); err != nil {
		return err
	}

	return resourceLayer0ServiceRead(d, meta)
}

func resourceLayer0ServiceDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(client.Client)
	serviceID := d.Id()

	jobID, err := client.DeleteService(serviceID)
	if err != nil {
		if strings.Contains(err.Error(), "No service found") {
			return nil
		}

		return err
	}

	if err := client.WaitForJob(jobID, defaultTimeout); err != nil {
		return err
	}

	return nil
}