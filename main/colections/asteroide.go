components {
  id: "rocha"
  component: "/main/scripts/rocha.script"
}
embedded_components {
  id: "model"
  type: "model"
  data: "mesh: \"/main/models/asteroid.glb\"\n"
  "name: \"{{NAME}}\"\n"
  "materials {\n"
  "  name: \"Realistic_Red_Asteroid_mat\"\n"
  "  material: \"/builtins/materials/model.material\"\n"
  "  textures {\n"
  "    sampler: \"tex0\"\n"
  "    texture: \"/main/models/asteroid.png\"\n"
  "  }\n"
  "}\n"
  ""
}
