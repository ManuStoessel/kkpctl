package cmd

import (
	"fmt"
	"io/ioutil"

	"github.com/cedi/kkpctl/pkg/client"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	writeConfig bool
)

// clustersCmd represents the clusters command
var getKubeconfigCmd = &cobra.Command{
	Use:               "kubeconfig [clusterid]",
	Short:             "Lists clusters for a given project (and optional seed datacenter) or fetch a named cluster.",
	Long:              `If no clusterid is specified, all clusters of an project are listed. If a clusterid is specified only this cluster is shown`,
	Example:           "kkpctl get kubeconfig --project x5zvx9bcx6 --datacenter es1",
	Args:              cobra.MaximumNArgs(1),
	ValidArgsFunction: getValidClusterArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		//fmt.Println("clusters called")
		kkp, err := client.NewClient(baseURL, apiToken)
		if err != nil {
			return errors.Wrap(err, "Could not initialize Kubermatic API client")
		}

		result, err := kkp.GetKubeConfig(args[0], projectID, datacenter)
		if err != nil {
			return errors.Wrap(err, "Error fetching kubeconfig")
		}

		if writeConfig == false {
			fmt.Println(result)
		} else {
			err := ioutil.WriteFile(fmt.Sprintf("kubeconfig-admin-%s", args[0]), []byte(result), 0644)
			if err != nil {
				return errors.Wrap(err, "Error writing kubeconfig")
			}
		}

		return nil
	},
}

func init() {
	getCmd.AddCommand(getKubeconfigCmd)

	getKubeconfigCmd.Flags().StringVarP(&projectID, "project", "p", "", "ID of the project.")
	getKubeconfigCmd.MarkFlagRequired("project")

	getKubeconfigCmd.Flags().StringVarP(&datacenter, "datacenter", "d", "", "Name of the datacenter.")
	getKubeconfigCmd.MarkFlagRequired("datacenter")

	getKubeconfigCmd.Flags().BoolVarP(&writeConfig, "write", "w", false, "write the kubeconfig to the local directory")
}