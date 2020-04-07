package sonar

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// Workflow to unmarshal to
type Workflow struct {
	Name      string
	CreatedAt string
}

func dataSourceSonarWorkflow() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSonarWorkflowRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Description: "Name of workflow",
				Required:    true,
			},
			"created_at": {
				Type:        schema.TypeString,
				Description: "Created at",
				Computed:    true,
			},
		},
	}
}

func dataSourceSonarWorkflowRead(d *schema.ResourceData, meta interface{}) error {
	workflow := new(Workflow)

	name := d.Get("name").(string)
	client := meta.(*Client)

	res, err := client.DoRequest("GET", fmt.Sprintf("/workflows/%s", name), nil, nil)
	if err != nil {
		log.Panicln(err)
	}

	defer res.Body.Close()
	json.NewDecoder(res.Body).Decode(&workflow)
	log.Printf("name: %s", workflow.Name)

	d.SetId(workflow.Name)
	d.Set("name", workflow.Name)
	d.Set("created_at", workflow.CreatedAt)

	return nil
}
