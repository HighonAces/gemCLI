package main

//
//import (
//	"context"
//	"log"
//	"os"
//
//	"github.com/google/generative-ai-go/genai"
//	"google.golang.org/api/option"
//)
//
//func GenerateTextFromAPI(ctx context.Context, text string) (*genai.Content, error) {
//	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("API_KEY")))
//	if err != nil {
//		return nil, err
//	}
//	defer client.Close()
//
//	model := client.GenerativeModel("gemini-2.0-flash-exp")
//	model.SetTemperature(1)
//	model.SetTopK(40)
//	model.SetTopP(0.95)
//	model.SetMaxOutputTokens(8192)
//	model.ResponseMIMEType = "application/json"
//	model.SystemInstruction = &genai.Content{
//		Parts: []genai.Part{genai.Text("command-line style with 7 examples")},
//	}
//
//	parts := []genai.Part{
//		genai.Text("input: echo command usage"),
//		genai.Text("output: 👉 Printing a String\n➡️ echo Hello, World!\n\n👉 Writing to a File\n➡️ echo -e 'Hello, World! \\nThis is PNAP!' >> test.txt\n\n👉 Writing to a File and the Terminal\n➡️ echo \"Hello, world!\" | tee output.txt\n\n👉Displaying a Variable Value\n➡️ echo $USER\n\n👉Omitting Trailing Newline\n➡️ echo -n \"Enter your name: \"\n\n👉Listing the Current Directory Contents\n➡️ echo */\n\n👉Listing Specific Types of Files\n➡️ echo *.txt"),
//		genai.Text("input: ping command usage"),
//		genai.Text("output: 👉 ping Command Syntax\n➡️ ping [options] [hostname/IP address]\n\n👉 Specify the Internet Protocol\n➡️ ping -6 [hostname]\n\n👉Change Time Interval Between Ping Packets\n➡️ ping -i 0.5 google.com\n\n👉Limit Number of Ping Packets\n➡️ ping -c 2 google.com\n\n👉Suppress ping Output to Print only Summary Statistics\n➡️ ping -c 10 -q google.com"),
//		genai.Text("input: command to exec into pod"),
//		genai.Text("output: 👉 Basic Command Execution in a Pod\n➡️ kubectl exec -it"),
//		genai.Text("input: nslookup command"),
//		genai.Text("output: 👉 Basic DNS Lookup\n➡️ nslookup google.com\n\n👉 Specify a Different DNS Server\n➡️ nslookup google.com 8.8.8.8\n\n👉 Query for Specific Record Type\n➡️ nslookup -type=MX google.com\n\n👉 Interactive Mode\n➡️ nslookup\n   > set type=A\n   > google.com\n   > exit\n\n👉 Display Debugging Information\n➡️ nslookup -debug google.com\n\n👉View Domain's NS Records\n➡️ nslookup -type=ns [domain-name]\n\n👉 View Information About a Specific Name Server\n➡️ nslookup [domain-name] [name-server]"),
//		genai.Text("input: command to search text"),
//		genai.Text("output: 👉 Basic Text Search in a File\n➡️ grep \"search_term\" file.txt\n\n👉 Case-Insensitive Text Search\n➡️ grep -i \"search_term\" file.txt\n\n👉 Recursive Search in Directories\n➡️ grep -r \"search_term\" directory/\n\n👉 Search for a Whole Word\n➡️ grep -w \"search_term\" file.txt\n\n👉 Display Line Numbers With Results\n➡️ grep -n \"search_term\" file.txt\n\n👉 Display Only Matching File Names\n➡️ grep -l \"search_term\" *.txt\n\n👉 Invert the Search\n➡️ grep -v \"search_term\" file.txt\n\n👉 Search Multiple Files\n➡️ grep \"search_term\" file1.txt file2.txt"),
//		genai.Text("input: use jq to get name and ip from pod"),
//		genai.Text("output: 👉 Extracting Name and IP from Pod JSON\n➡️ kubectl get pods -o json | jq '.items[] | {name: .metadata.name, ip: .status.podIP}'\n\n‼️ Other example usage \n👉 list pod status\n➡️ kubectl -n"),
//		genai.Text("input: helm commands"),
//		genai.Text("output: 👉 Installing a Helm Chart\n➡️ helm install my-release stable/my-chart\n\n👉 Listing Installed Helm Releases\n➡️ helm list\n\n👉 Uninstalling a Helm Release\n➡️ helm uninstall my-release\n\n👉 Upgrading a Helm Release\n➡️ helm upgrade my-release stable/my-chart\n\n👉 Getting Helm Chart Information\n➡️ helm show chart stable/my-chart\n\n👉 Getting Helm Chart Values\n➡️ helm show values stable/my-chart\n\n👉 Creating a New Chart\n➡️ helm create my-chart\n\n👉 Fetching a Helm Chart\n➡️ helm fetch stable/my-chart"),
//		genai.Text("input: input:"),
//		genai.Text("output: output:"),
//		genai.Text("input: echo command usage"),
//		genai.Text("output: 👉 Printing a String\n➡️ echo Hello, World!\n\n👉 Writing to a File\n➡️ echo -e 'Hello, World! \\nThis is PNAP!' >> test.txt\n\n👉 Writing to a File and the Terminal\n➡️ echo \"Hello, world!\" | tee output.txt\n\n👉Displaying a Variable Value\n➡️ echo $USER\n\n👉Omitting Trailing Newline\n➡️ echo -n \"Enter your name: \"\n\n👉Listing the Current Directory Contents\n➡️ echo */\n\n👉Listing Specific Types of Files\n➡️ echo *.txt"),
//		genai.Text("input: ping command usage"),
//		genai.Text("output: 👉 ping Command Syntax\n➡️ ping [options] [hostname/IP address]\n\n👉 Specify the Internet Protocol\n➡️ ping -6 [hostname]\n\n👉Change Time Interval Between Ping Packets\n➡️ ping -i 0.5 google.com\n\n👉Limit Number of Ping Packets\n➡️ ping -c 2 google.com\n\n👉Suppress ping Output to Print only Summary Statistics\n➡️ ping -c 10 -q google.com"),
//		genai.Text("input: nslookup command"),
//		genai.Text("output: 👉 Basic DNS Lookup\n➡️ nslookup google.com\n\n👉 Specify a Different DNS Server\n➡️ nslookup google.com 8.8.8.8\n\n👉 Query for Specific Record Type\n➡️ nslookup -type=MX google.com\n\n👉 Interactive Mode\n➡️ nslookup\n   > set type=A\n   > google.com\n   > exit\n\n👉 Display Debugging Information\n➡️ nslookup -debug google.com\n\n👉View Domain's NS Records\n➡️ nslookup -type=ns [domain-name]\n\n👉 View Information About a Specific Name Server\n➡️ nslookup [domain-name] [name-server]"),
//		genai.Text("input: command to search text"),
//		genai.Text("output: 👉 Basic Text Search in a File\n➡️ grep \"search_term\" file.txt\n\n👉 Case-Insensitive Text Search\n➡️ grep -i \"search_term\" file.txt\n\n👉 Recursive Search in Directories\n➡️ grep -r \"search_term\" directory/\n\n👉 Search for a Whole Word\n➡️ grep -w \"search_term\" file.txt\n\n👉 Display Line Numbers With Results\n➡️ grep -n \"search_term\" file.txt\n\n👉 Display Only Matching File Names\n➡️ grep -l \"search_term\" *.txt\n\n👉 Invert the Search\n➡️ grep -v \"search_term\" file.txt\n\n👉 Search Multiple Files\n➡️ grep \"search_term\" file1.txt file2.txt"),
//		genai.Text("input: use jq to get name and ip from pod"),
//		genai.Text("output: 👉 Extracting Name and IP from Pod JSON\n➡️ kubectl get pods -o json | jq '.items[] | {name: .metadata.name, ip: .status.podIP}'\n\n‼️ Other example usage \n👉 list pod status\n➡️ kubectl -n"),
//		genai.Text("input: command to exec into pod"),
//		genai.Text("output: 👉 Basic Command Execution in a Pod\n➡️ kubectl exec -it"),
//		genai.Text("input: helm commands"),
//		genai.Text("output: 👉 Installing a Helm Chart\n➡️ helm install my-release stable/my-chart\n\n👉 Listing Installed Helm Releases\n➡️ helm list\n\n👉 Uninstalling a Helm Release\n➡️ helm uninstall my-release\n\n👉 Upgrading a Helm Release\n➡️ helm upgrade my-release stable/my-chart\n\n👉 Getting Helm Chart Information\n➡️ helm show chart stable/my-chart\n\n👉 Getting Helm Chart Values\n➡️ helm show values stable/my-chart\n\n👉 Creating a New Chart\n➡️ helm create my-chart\n\n👉 Fetching a Helm Chart\n➡️ helm fetch stable/my-chart"),
//		genai.Text("input: find command with xargs"),
//		genai.Text("output: 👉 Finding Files and Executing Commands with xargs\n➡️ find . -name \"*.txt\" -print0 | xargs -0 grep \"search_term\"\n\n👉 Finding Files and Deleting Them\n➡️ find . -name \"*.log\" -print0 | xargs -0 rm\n\n👉 Finding Files and Changing Permissions\n➡️ find . -type f -print0 | xargs -0 chmod 644\n\n👉 Finding Directories and Listing Them\n➡️ find . -type d -print0 | xargs -0 ls -ld\n\n👉 Running a Command on Found Files in Parallel\n➡️ find . -name \"*.jpg\" -print0 | xargs -0 -P 4 convert -resize 800x600\n\n👉 Executing a Command that Needs Multiple Arguments\n➡️ find . -type f -name \"*.mp4\" -print0 | xargs -0 -I {} ffmpeg -i {} -c:v libx264 {}.mp4input: scp command usageoutput: 👉 Securely Copying a Local File to a Remote Server\n➡️ scp local_file.txt user@remote_host:/path/to/destination\n\n👉 Securely Copying a Local Directory to a Remote Server\n➡️ scp -r local_directory user@remote_host:/path/to/destination\n\n👉 Securely Copying a Remote File to Local Machine\n➡️ scp user@remote_host:/path/to/remote_file.txt local_destination\n\n👉 Securely Copying a Remote Directory to Local Machine\n➡️ scp -r user@remote_host:/path/to/remote_directory local_destination\n\n👉 Using a Different Port for SCP\n➡️ scp -P 2222 local_file.txt user@remote_host:/path/to/destination\n\n👉 Preserving Modification Times and Access Times\n➡️ scp -p local_file.txt user@remote_host:/path/to/destination\n\n👉 Limiting Bandwidth Usage\n➡️ scp -l 100 local_file.txt user@remote_host:/path/to/destination"),
//		genai.Text("input: aws eks commands"),
//		genai.Text("output: 👉 Creating an EKS Cluster\n➡️ aws eks create-cluster --name my-cluster --region us-west-2 --version 1.23 --role-arn arn:aws:iam::123456789012:role/eks-cluster-role --resources-vpc-config subnetIds=subnet-0123456789abcdef0,subnet-abcdef01234567890,securityGroupIds=sg-0123456789abcdef0\n\n👉 Listing EKS Clusters\n➡️ aws eks list-clusters\n\n👉 Describing an EKS Cluster\n➡️ aws eks describe-cluster --name my-cluster\n\n👉 Updating an EKS Cluster\n➡️ aws eks update-cluster-version --name my-cluster --version 1.24\n\n👉 Deleting an EKS Cluster\n➡️ aws eks delete-cluster --name my-cluster\n\n👉 Creating an EKS Node Group\n➡️ aws eks create-nodegroup --cluster-name my-cluster --nodegroup-name my-nodegroup --node-role arn:aws:iam::123456789012:role/eks-node-role --subnets subnet-0123456789abcdef0,subnet-abcdef01234567890 --scaling-config minSize=2,maxSize=4,desiredSize=2 --instance-types t3.medium\n\n👉 Updating an EKS Node Group\n➡️ aws eks update-nodegroup-version --cluster-name my-cluster --nodegroup-name my-nodegroup --version 1.24\n\n👉 Deleting an EKS Node Group\n➡️ aws eks delete-nodegroup --cluster-name my-cluster --nodegroup-name my-nodegroup\n\n👉 Getting Credentials to Access Cluster\n➡️ aws eks --region us-west-2 update-kubeconfig --name my-cluster"),
//		genai.Text("input: "),
//		genai.Text("output: "),
//	}
//
//	resp, err := model.GenerateContent(ctx, parts...)
//	if err != nil {
//		log.Fatalf("Error sending message: %v", err)
//	}
//
//	return resp.Candidates[0].Content, nil
//
//}
