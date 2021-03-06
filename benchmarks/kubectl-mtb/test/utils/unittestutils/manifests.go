package unittestutils

var ClusterPolicy = `
apiVersion: kyverno.io/v1
kind: ClusterPolicy
metadata:
  name: disallow-privileged
spec:
  validationFailureAction: enforce
  rules:
    - name: validate-privileged
      match:
        resources:
          kinds:
            - Pod
          namespaces:
            - tenant1admin
      validate:
        message: "Privileged mode is not allowed. Set privileged to false"
        pattern:
          spec:
            containers:
              - =(securityContext):
                  # https://github.com/kubernetes/api/blob/7dc09db16fb8ff2eee16c65dc066c85ab3abb7ce/core/v1/types.go#L5707-L5711
                  # k8s default to false
                  =(privileged): false
`
