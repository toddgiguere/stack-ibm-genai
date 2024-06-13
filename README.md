# WatsonX GenAI on the IBM Cloud

The following [deployable architecture](https://cloud.ibm.com/docs/secure-enterprise?topic=secure-enterprise-understand-module-da#what-is-da) automates the deployment of a GenAI Pattern on IBM Cloud, including all underlying IBM Cloud infrastructure. This architecture implements the best practices for Watsonx GenAI Pattern deployment on IBM Cloud.

This deployable architecture provides a comprehensive foundation for trust, observability, security, and regulatory compliance by configuring the IBM Cloud account to align with compliance settings, deploying key and secret management services, and deploying the infrastructure to support CI/CD/CC pipelines for secure application lifecycle management. These pipelines facilitate the deployment of the application, vulnerability checks, and auditability, ensuring a secure and trustworthy deployment of Generative AI applications on IBM Cloud.

# Objective and Benefits

This deployable architecture is designed to showcase a fully automated deployment of a WatsonX GenAI through IBM Cloud Project, providing a flexible and customizable foundation for your own Watson-based application deployments on IBM Cloud.

By leveraging this architecture, you can accelerate your deployment and tailor it to meet your unique business needs and enterprise goals.

By using this architecture, you can:

* Establish Trust: The architecture ensures trust by configuring the IBM Cloud account to align with compliance settings as defined in the [Financial Services](https://cloud.ibm.com/docs/framework-financial-services?topic=framework-financial-services-about) framework.
* Ensure Observability: The architecture provides observability by deploying services such as IBM Log Analysis, IBM Monitoring, IBM Activity Tracker, and log retention through Cloud Object Storage buckets.
* Implement Security: The architecture ensures security by deploying IBM Key Protect and IBM Secrets Manager.
* Achieve Regulatory Compliance: The architecture ensures regulatory compliance with IBM Security Compliance Center (SCC).


# Deployment Details

To deploy this architecture, follow these steps.

## 1. Prerequisites

Before deploying the deployable architecture, ensure you have:

* Created an API key in the target account with sufficient permissions. The target account is the account that will be hosting the resources deployed by this deployable architecture. See [instructions](https://cloud.ibm.com/docs/account?topic=account-userapikey&interface=ui) Note the API key, as it will be used later. On evaluation environments, you may simply grant `Administrator` role on `IAM Identity Service`, `All Identity and Access enabled services` and `All Account Management` services. If you need to narrow down further access, for a production environment for instance, the minimum level of permissions is indicated in the [Permission tab](https://cloud.ibm.com/catalog/7a4d68b4-cf8b-40cd-a3d1-f49aff526eb3/architecture/Retrieval_Augmented_Generation_Pattern-5fdd0045-30fc-4013-a8bc-6db9d5447a52-global#permissions) of the deployable architecture.
* (Optional) Installed the IBM Cloud CLI's Project add-on using the `ibmcloud plugin install project` command. More information is available [here](https://cloud.ibm.com/docs/cli?topic=cli-projects-cli).

Ensure that you are familiar with the "Important Deployment Considerations" located at the bottom of this document.

## 2. Deploy the Stack in a New Project from Catalog

* Locate the [tile](https://cloud.ibm.com/catalog/7a4d68b4-cf8b-40cd-a3d1-f49aff526eb3/architecture/Retrieval_Augmented_Generation_Pattern-5fdd0045-30fc-4013-a8bc-6db9d5447a52-global) for the Deployable Architecture in the IBM Cloud Catalog.
* Click the "Add to project" button.

  ![image](./images/min/1-catalog.png)

* Select **Create new** and enter the following details:
    - Name and Description (e.g., "WatsonX GenAI Pattern")
    - Region and Resource Group for the project. e.g. for evaluation purposes, you may select the region the closest to you, and the Default resource group. For more insights on the recommended production topology, refer to the Enterprise account architecture Central administration account [white paper](https://cloud.ibm.com/docs/enterprise-account-architecture?topic=enterprise-account-architecture-admin-hub-account).
    - Configuration Name (name of the automation in the project, e.g., "genai", "dev" or "prod", ideally matching the deployment target, but this can be any name)

      ![project](./images/min/2-project.png)

* Click the **Add** button (or **Create** if this is the first project in the account) at the bottom right of the modal popup to complete.

## 3. Set the Input Configuration for the Stack

After completing `Step 2 - Deploy the Stack in a New Project from Catalog`, you are directed to a page allowing you to enter the configuration for you deployment:
* Under Security -> Authentication, enter the API Key from the prereqs in the `api_key` field.
  ![inputs](./images/min/3-inputs.png)
* Under Required, input a prefix. This prefix will be appended to the name of most resources created by automation, ensuring uniqueness and avoiding clashes when provisioning names in the same account.
* Under Optional, input the signing_key field. While not necessary for deploying Cloud resources, it is recommended and required to enable the building and deployment of the sample app.

You may explore the other available inputs, such as the region and resource group name (under optional tab), leave them as is, or modify them as needed.

Once ready, click the "Save" button at the top of the screen.

## 4. Deploy the Architecture

Navigate to the project deployment view by clicking the project name in the breadcrumb menu.

![menu](./images/min/4-bread.png)


You should be directed to a screen looking like:

![validate](./images/min/5-validate.png)

Note: in some rare occurrences, the first member of the stack may not be marked as "Ready to validate". Refreshing the page in your browser window should solve this problem.

Approach to deploy the architecture:
1. Through the UI

### Approach: Deployment through the UI

1. Click on validate

   ![validate button](./images/min/5b-validate.png)

2. Wait for validation

   ![validation](./images/min/6-validation.png)

3. Approve and click the deploy button

   ![deploy](./images/min/7-deploy.png)

4. Wait for deployment

5. Repeat step 1 for the next configuration in the architecture. Note that as you progress in deploying the initial base configuration, you will be given the option to validate and deploy multiple configuration in parallel.

## 5. Important Deployment Considerations

### API Key Requirements

The deployable architecture can only be deployed with an API Key associated with a user. It is not compatible with API Keys associated with a serviceId. Additionally, it cannot be deployed using the Project trusted profile support.

### Known UI Issue: "Unable to validate your configuration"

After approving the configuration, you may encounter an error message stating "Unable to validate your configuration". This is a known UI issue that can be resolved by simply **refreshing your browser window**. This will allow you to continue with the deployment process.

### Notification of New Configuration Versions ("Needs Attention")

You may see notifications in IBM Cloud Project indicating that one or more configurations in the stack have new versions available. You can safely ignore these messages at this point, as they will not prevent you from deploying the stack. No specific action is required from you.

![new version](./images/min/10-new-version.png)

Please note that these notifications are expected, as we are rapidly iterating on the development of the underlying components. As new stack versions become available, the versions of the underlying components will also be updated accordingly.

### Limitations with the Trial Secret Manager Offering

The automation is configured to deploy a Trial version of Secret Manager by default to minimize costs. However, the Trial version has some limitations. If you want to avoid these limitations, you can opt to deploy a standard (paid) instance of Secret Manager under the **Optional settings** of the stack.

Here are the limitations of the Trial version:
* **Account limitation**: Only one Trial instance of Secret Manager can be deployed at a time in a given account.
* **Deployment error**: You will encounter an error in the Secret Manager deployment step if there is already a Trial instance deployed in the same account.
* **Re-deployment failure**: If the automation provisions a Trial version of Secrets Manager, and is un-deployed and then re-deployed again with the Trial version in the same account, the "2b - Security Service - Secret Manager" deployment will fail. This is because you can only have one Trial version of Secrets Manager in an account, and even after deletion, the prior Trial version of Secrets Manager needs to be removed from the "reclamation" state as well.


**What are reclamations?**
In IBM Cloud, when you delete a resource, it doesn't immediately disappear. Instead, it enters a "reclamation" state, where it remains for a short period of time (usually 7 days) before being permanently deleted. During this time, you can still recover the resource if needed.

To resolve the re-deployment failure, you will need to delete the Secret Manager service from the reclamation state by running the following commands:
```
ibmcloud resource reclamations #  lists all the resources in reclamation state, get the reclamation ID of the secret manager service
ibmcloud resource reclamation-delete <reclamation-id>
```


# Customization options

There are numerous customization possibilities available out of the box. This section explores some common scenarios, but is not exhaustive.

## Editing Individual Configurations

Each configuration in the deployed stack surfaces a large number of input parameters. You can directly edit each parameter to tailor your deployment by selecting the **Edit** option in the menu for the corresponding configuration on the right-hand side.

![edit config](./images/min/11-edit-config.png)

This approach enables you to:
- Fine-tune account settings
- Deploying additional Watson components, such as Watsonx Governance
- Deploy to an existing resource group
- Reuse existing key protect keys
- Tuning the parameter of the provisioned code engine project
- ...

## Removing Configurations from the Stack

You can remove any configuration from the stack, provided there is no direct dependency in later configurations, by selecting the **Remove from Stack** option in the right-hand side menu for the corresponding configuration.

This applies to the following configurations:
- Observability
- Security and Control Center

![edit config](./images/min/12-remove-config.png)

## Managing Stack-Level Inputs and Outputs

You can add or remove inputs and outputs surfaced at the stack level by following these steps:
1. Select the stack configuration

   ![stack def](./images/min/13-define-stack.png)
1. You are presented with a screen allowing you to promote any of the configuration inputs or outputs at the stack level

   ![stack def](./images/min/14-stack-def.png)


## Sharing Modified Stacks through a Private IBM Cloud Catalog

Once you have made modifications to your stack in Project, you can share it with others through a private IBM Cloud Catalog. To do so, follow these steps:
1. Deploy the stack at least once: You need to deploy the stack first to allow importing the stack definition to a private catalog.
2. Select the "Add to private catalog" option in the menu located on the stack configuration.

This will allow you to share your modified stack with others through a private IBM Cloud Catalog.

# Undeploying/Deleting the Stack, and all associated Infrastructure Resources

## Undeploying Infrastructure

To undeploy the infrastructure created by the automation, complete the following steps:

### 1. Undeploy Configurations in the Project

Undeploy each configuration in the project, one by one, via UI, starting from the "4 - SaaS DA" and working your way up in the stack up to, and inclusive of "2a - Security Service - Key Management". Wait for full undeployment of a configuration before starting to undeploy the next configuration up in the stack.

### 2. Delete Reclamation Claims

Before undeploying the "1 - Account Infrastructure Base", you will need to manually delete the reclamation claims for the resources deleted from the previous steps. Reclamation allows you to restore deleted resources for up to one week. However, any reclamation that is still active prevents from deleting the resource group managed by the "1 - Account Infrastructure Base":
* Log in to the target IBM Cloud account with the CLI
* Run `ibmcloud resource reclamations` to view the full list of reclamation. You may identify the exact reclamations to delete as they are planned to be deleted in one week after the date for which the resource was deleted.
* For each reclamation, execute `ibmcloud resource reclamation-delete <reclamation-id>`. The reclamation-id is the id provided in the results from ibmcloud resource reclamations listing.
* Run `ibmcloud resource reclamations` again to ensure the reclamations have been fully deleted

More details are available [here](https://cloud.ibm.com/docs/account?topic=account-resource-reclamation&interface=cli).

### 3. Undeploy "1 - Account Infrastructure Base"

You may now undeploy "1 - Account Infrastructure Base" in the project.

### 4. Delete Project

Once all configurations are undeployed, you may delete the project.
