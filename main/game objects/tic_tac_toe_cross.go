components {
  id: "player_object"
  component: "/main/scripts/games/tic tac toe/player_object.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"Cross\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "size {\n"
  "  x: 160.0\n"
  "  y: 160.0\n"
  "}\n"
  "size_mode: SIZE_MODE_MANUAL\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/atlases/tic_tac_toe.atlas\"\n"
  "}\n"
  ""
  scale {
    x: 0.8
    y: 0.8
  }
}
