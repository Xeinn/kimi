package main

import (
  "sigs.k8s.io/kustomize/api/resmap"
)

type plugin struct {
  someValue string
}

var KustomizePlugin plugin

func (p *plugin) Config(h *resmap.PluginHelpers, c []byte) error {

  return nil
}

func (p *plugin) Generate() (resmap.ResMap, error) {

  return nil, nil
}

func (p *plugin) Transform(m resmap.ResMap) error {

  return nil
}
