# Elasticsearch + Kibana + Filebeat on Amazon EKS

This project sets up an Elasticsearch logging stack on Kubernetes, consisting of:

- **Elasticsearch** (stateful search and analytics engine)
- **Kibana** (visual dashboard for Elasticsearch)
- **Filebeat** (log shipper to send pod logs to Elasticsearch)

All resources are deployed under the `aa-upwork` namespace.

---

## 🧱 Prerequisites

- AWS CLI configured
- eksctl installed
- kubectl installed
- Go (for script execution)
- Basic knowledge of YAML and Kubernetes concepts

---

## 📁 Project Structure

```bash
k8s-elasticsearch
├── elasticsearch.yaml       # Deploys Elasticsearch StatefulSet + Service
├── kibana.yaml              # Deploys Kibana Deployment + Services + ConfigMap
├── filebeat.yaml            # Deploys Filebeat DaemonSet
└── helper.go                # Go script to dynamically inject Elasticsearch pod IP into Kibana config

```
## 🚀 Setup Instructions

1. **EKS Cluster Setup:** </br>
	Create an EKS cluster using eksctl: ```eksctl create cluster --name K8s-cluster --region us-east-1 --nodes 2 --node-type t3.medium```
2. **Create the Namespace:** ```kubectl create namespace aa-upwork```
3. **Deploy Elasticsearch:** ```kubectl create -f elastic.yaml``` </br>
   Wait for the pod to be ready: ```kubectl get pods -n aa-upwork```
4. **Update Kibana Config with Elasticsearch Pod IP** </br>
    Run the Go script to inject the pod IP into the Kibana config: ```go run helper.go kibana.yaml```
   > This will modify 2-kibana.yaml in-place, replacing <IP> with the actual IP of the elasticsearch-0 pod.
5. **Deploy Kibana:** ```kubectl create -f kibana.yaml```
6. **Deploy Filebeat:** ```kubectl create -f filebeat.yaml```

## 🌐 Access Kibana

Once the Kibana service is running, find the external IP: ```kubectl get svc -n aa-upwork``` </br>
Look for the service named kibana with EXTERNAL-IP and open: ```http://<EXTERNAL-IP>:5601``` </br>

**For Example -**

<img width="1093" height="103" alt="Screenshot 2025-07-29 at 4 28 35 PM" src="https://github.com/user-attachments/assets/f402fcf7-6ca2-45fc-892d-22d1b681f07e" />

we can access kibana using URL ```http://a77cddf2ca6a24db3bc644cc86a85e95-417137951.us-east-1.elb.amazonaws.com:5601/``` </br>

<img width="1470" height="956" alt="Screenshot 2025-07-29 at 4 16 03 PM" src="https://github.com/user-attachments/assets/34d462c0-8c0b-4ca4-b307-3c959faaf6aa" />

## 🔧 Filebeat Integration in Kibana

Once Filebeat is deployed, you can set up its integration inside Kibana to view logs from all pods across your cluster:</br>

1. Open Kibana in your browser using: ```http://<EXTERNAL-IP>:5601```
2. Go to “Stack Management” > “Index Patterns” (or “Data Views” in newer Kibana versions).
3. Click on “Create index pattern”, and enter: ```filebeat-*```
4. Set the timestamp field to @timestamp, then click “Create”.
5. Now go to “Discover”. You should see logs coming in from your Kubernetes pods via Filebeat.


## 🔁 Cleanup
To delete cluster (Created via eksctl): ```eksctl delete cluster --name <cluster-name> --region <region>```

## 📌 Notes
	•	update_kibana_ip.go is a helper script written in Go to auto-patch the Elasticsearch IP into the Kibana config.
	•	Filebeat is configured to collect logs from all running pods in the cluster and forward them to Elasticsearch.

 ## ✍️ Author
  Shivam Singh Maurya
 
