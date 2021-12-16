#!/usr/bin/env python3

from diagrams import Cluster, Diagram
from diagrams.k8s.compute import Pod
from diagrams.k8s.rbac import ServiceAccount
from diagrams.k8s.others import CRD

with Diagram("Tig operator", show=False):
    with Cluster("tig-operator-system"):
        controller_pod = Pod("pod")

    with Cluster("namespace"):
        service_account = ServiceAccount("it")
        new_object = CRD("tig")

    service_account << controller_pod
    new_object << controller_pod

