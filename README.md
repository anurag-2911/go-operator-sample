# Time Scaler Operator

This project contains a Kubernetes Operator built with the Operator SDK and Go. 
It's designed to scale deployments based on time schedules.


As Operator-SDK needs a Linux operating system, so below are the steps for developing operators on Github Codespace
which has Ubuntu OS setup:

The steps are from the official documentation:
https://sdk.operatorframework.io/docs/building-operators/golang/quickstart/

# Check the OS architecture and Name, for example, for x86_64 architecture for an Ubuntu OS.
for version v1.26.1 of operator SDK:

curl -LO https://github.com/operator-framework/operator-sdk/releases/download/v1.26.1/operator-sdk_linux_amd64

# Make the binary executable

chmod +x operator-sdk_linux_amd64

# Move the binary to a directory in your PATH

sudo mv operator-sdk_linux_amd64 /usr/local/bin/operator-sdk

# Verify the installation
operator-sdk version

# Initialize the Operator Project
operator-sdk init --domain example.com --repo github.com/anurag-2911/go-operator-sample

Create an API and Controller ( this will generate the starter files and code)

operator-sdk create api --group autoscaling --version v1 --kind TimeScaler --resource --controller

The above commands will create files
a. controllers/timescaler_controller.go
b. api/v1/timescaler_types.go
c. main.go
d. Makefile

# Deploying this controller and CRD (Custom Resource Definition) to a Kubernetes cluster involves several steps,
including building the operator's image, pushing it to a container registry, and applying the necessary Kubernetes manifests

Deploying your controller and CRD (Custom Resource Definition) to a Linode Kubernetes cluster involves several steps, including building your operator's image, pushing it to a container registry, and applying the necessary Kubernetes manifests. Here's a step-by-step guide:

### Step 1: Build the Operator Docker Image

1. For example, your-registry/time-scaler-operator:v1.0`,
where `your-registry` is the Docker Hub username or the URL of another container registry.

2. **Build the Image**: Build the Docker image from the root of your operator project.
Make sure Docker is running on your machine.

operator-sdk build your-registry/time-scaler-operator:v1.0  

   
### Step 2: Push the Docker Image to a Registry

1. **Log In to Your Container Registry**: Ensure you're logged into the registry where you're pushing the image.
   
docker login
   

2. **Push the Image**: Push your operator's Docker image to the registry.
   
   docker push your-registry/time-scaler-operator:v1.0
   

### Step 3: Update Your Operator's Deployment Manifest

1. **Edit the Operator's Deployment Manifest**: Find the deployment manifest for your operator.
It's located in the `config/manager/manager.yaml` file or a similar path within your project.

2. **Set the Image**: Update the `image` field under `containers` to match the image pushed.

   ```yaml
   spec:
     containers:
     - name: manager
       image: your-registry/time-scaler-operator:v1.0
       ...
   ```

### Step 4: Deploy the CRD to the Cluster

1. **Apply the CRD Manifests**: From the root of the operator project,
apply the CRDs to the Kubernetes cluster. Ensure the `kubectl` context is set to the Kubernetes cluster.

   
   kubectl apply -f config/crd/bases
 

### Step 5: Deploy the Operator

1. **Apply the Operator Manifests**: Now, deploy the operator to the cluster.
This will usually involve applying the manifests in the `config` directory.

   
   kubectl apply -f config/manager/manager.yaml
   

   Or, if your project uses Kustomize, you might use:

   
   kubectl apply -k config/default
 

### Step 6: Deploy a Custom Resource to Test the Operator

1. **Create a Custom Resource (CR)**: To test your operator, you'll need to deploy a CR that it manages.
Create a YAML file (e.g., `time-scaler-cr.yaml`) with your `TimeScaler` resource definition.

   ```yaml
   apiVersion: autoscaling.example.com/v1
   kind: TimeScaler
   metadata:
     name: example-timescaler
   spec:
     schedule: "*/5 * * * *"
     replicas: 2
   ```

2. **Apply the CR**: Apply this configuration to your Linode Kubernetes cluster.

   ```sh
   kubectl apply -f time-scaler-cr.yaml
   ```

### Step 7: Verify Everything is Working

1. **Check the Operator Logs**: You can check the logs of your operator to ensure it's working as expected.

   
   kubectl logs deployment/<your-operator-deployment-name> -n <your-operator-namespace>
   

2. **Verify the Scaling Action**: Confirm that the scaling action defined in your CR is being performed at the scheduled times.

By following these steps, it will have deployed the custom Kubernetes operator to the Kubernetes cluster,
and it should be operational, managing resources according to the above logic implemented.







