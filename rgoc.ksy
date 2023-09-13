meta:
  id: rgoc
  title: Ren'G Object Compiled
  file-extension: rgoc
  license: MIT
  encoding: UTF-8
  endian: le

seq:
  - id: magic
    contents: RENG

  - id: rgoc_version
    type: strz
    encoding: UTF-8

  - id: revision
    type: strz
    encoding: UTF-8

  - id: num_labels
    type: u4

  - id: labels
    type: label
    repeat: expr
    repeat-expr: num_labels

  - id: num_screens
    type: u4

  - id: screens
    type: screen
    repeat: expr
    repeat-expr: num_screens

  - id: num_code_specs
    type: u4

  - id: code_specs
    type: code_spec
    repeat: expr
    repeat-expr: num_code_specs

enums:
  label_obj_type:
    0: code
    1: show
    2: hide
    3: play_video
    4: stop_video
    5: play_music
    6: stop_music
    7: play_channel
    8: stop_channel
    9: say
    10: pause
    11: call
    12: jump

  screen_obj_type:
    0: code
    1: show
    2: play_video
    3: play_music
    4: stop_music
    5: play_channel
    6: stop_channel
    7: text
    8: timer
    9: key
    10: button
    11: bar

  special_transform_type:
    0: center
    1: x_center
    2: y_center
    3: axis_center

  anime_type:
    0: alpha
    1: rotate
    2: x_pos
    3: y_pos

types:
  uuid:
    seq:
      - id: body
        size: 16

  color:
    seq:
      - id: r
        type: u1
      - id: g
        type: u1
      - id: b
        type: u1
      - id: a
        type: u1

  vec2:
    seq:
      - id: x
        type: f4
      - id: y
        type: f4

  transform:
    seq:
      - id: pos
        type: vec2
      - id: size
        type: vec2
      - id: flip
        type: vec2
      - id: rotate
        type: f4
      - id: special_transform
        type: special_transform

  special_transform:
    seq:
      - id: type
        type: u2
        enum: special_transform_type

      - id: body
        type:
          switch-on: type
          cases:
            "special_transform_type::center": center
            "special_transform_type::x_center": x_center
            "special_transform_type::y_center": y_center
            "special_transform_type::axis_center": axis_center

  center:
    seq:
      - id: placeholder
        contents: [0]

  x_center:
    seq:
      - id: y_pos
        type: f4

  y_center:
    seq:
      - id: x_pos
        type: f4

  axis_center:
    seq:
      - id: axis
        type: vec2

  label:
    seq:
      - id: name
        type: strz
        encoding: UTF-8

      - id: num_obj
        type: u8

      - id: obj
        type: label_object
        repeat: expr
        repeat-expr: num_obj

  screen:
    seq:
      - id: name
        type: strz
        encoding: UTF-8

      - id: num_obj
        type: u8

      - id: obj
        type: screen_object
        repeat: expr
        repeat-expr: num_obj

  label_object:
    seq:
      - id: type
        type: u2
        enum: label_obj_type

      - id: body
        type:
          switch-on: type
          cases:
            "label_obj_type::code": code
            "label_obj_type::show": show
            "label_obj_type::hide": hide
            "label_obj_type::play_video": play_video
            "label_obj_type::stop_video": stop_video
            "label_obj_type::play_music": play_music
            "label_obj_type::stop_music": stop_music
            "label_obj_type::play_channel": play_channel
            "label_obj_type::stop_channel": stop_channel
            "label_obj_type::say": say
            "label_obj_type::pause": pause
            "label_obj_type::call": call
            "label_obj_type::jump": jump

  screen_object:
    seq:
      - id: type
        type: u2
        enum: screen_obj_type

      - id: body
        type:
          switch-on: type
          cases:
            "screen_obj_type::code": code
            "screen_obj_type::show": show
            "screen_obj_type::play_video": play_video
            "screen_obj_type::play_music": play_music
            "screen_obj_type::stop_music": stop_music
            "screen_obj_type::play_channel": play_channel
            "screen_obj_type::stop_channel": stop_channel
            "screen_obj_type::text": text
            "screen_obj_type::timer": timer
            "screen_obj_type::key": key
            "screen_obj_type::button": button
            "screen_obj_type::bar": bar

  code:
    seq:
      - id: ref
        type: uuid

  code_spec:
    seq:
      - id: uuid
        type: uuid

      - id: num_instructions
        type: u8

      - id: instructions
        type: u1
        repeat: expr
        repeat-expr: num_instructions

  show:
    seq:
      - id: name
        type: strz
        encoding: UTF-8

      - id: transform
        type: transform

      - id: num_anime
        type: u8

      - id: anime
        type: anime
        repeat: expr
        repeat-expr: num_anime

  anime:
    seq:
      - id: loop
        type: b1

      - id: type
        type: u2
        enum: anime_type

      - id: init_value
        type: f8

      - id: start_time
        type: f8

      - id: duration
        type: f8

      - id: curve
        type: code

      - id: end
        type: code

  hide:
    seq:
      - id: texture_index
        type: u8

      - id: anime
        type: anime

  play_video:
    seq:
      - id: name
        type: strz
        encoding: UTF-8

      - id: transform
        type: transform

      - id: loop
        type: b1

      - id: num_anime
        type: u8

      - id: anime
        type: anime
        repeat: expr
        repeat-expr: num_anime

  stop_video:
    seq:
      - id: name
        type: strz
        encoding: UTF-8

      - id: num_anime
        type: u8

      - id: anime
        type: anime
        repeat: expr
        repeat-expr: num_anime

  play_music:
    seq:
      - id: path
        type: strz
        encoding: UTF-8

      - id: loop
        type: b1

      - id: ms
        type: u8

  stop_music:
    seq:
      - id: ms
        type: u8

  play_channel:
    seq:
      - id: path
        type: strz
        encoding: UTF-8

      - id: chan_name
        type: strz
        encoding: UTF-8

      - id: ms
        type: u8

  stop_channel:
    seq:
      - id: chan_name
        type: strz
        encoding: UTF-8

      - id: ms
        type: u8

  say:
    seq:
      - id: character
        type: character

      - id: text
        type: strz
        encoding: UTF-8

  character:
    seq:
      - id: name
        type: strz
        encoding: UTF-8

  pause:
    seq:
      - id: time
        type: f8

  call:
    seq:
      - id: label_name
        type: strz
        encoding: UTF-8

  jump:
    seq:
      - id: label_name
        type: strz
        encoding: UTF-8

  text:
    seq:
      - id: text
        type: strz
        encoding: UTF-8

      - id: font_name
        type: strz
        encoding: UTF-8

      - id: transform
        type: transform

      - id: color
        type: color

      - id: typing_fx
        type: b1

  timer:
    seq:
      - id: time
        type: f8

      - id: do
        type: code

  key:
    seq:
      - id: down
        type: code

      - id: up
        type: code

  button:
    seq:
      - id: main_image_name
        type: strz
        encoding: UTF-8

      - id: hover_image_name
        type: strz
        encoding: UTF-8

      - id: transform
        type: transform

      - id: num_anime
        type: u8

      - id: anime
        type: anime
        repeat: expr
        repeat-expr: num_anime

      - id: down
        type: code

      - id: up
        type: code

      - id: hover
        type: code

      - id: un_hover
        type: code

  bar:
    seq:
      - id: frame_image_name
        type: strz
        encoding: UTF-8

      - id: cursor_image_name
        type: strz
        encoding: UTF-8

      - id: curser_hover_image_name
        type: strz
        encoding: UTF-8

      - id: gauge_image_name
        type: strz
        encoding: UTF-8

      - id: frame_image_transform
        type: transform

      - id: cursor_size
        type: vec2

      - id: start_padding
        type: f4

      - id: end_padding
        type: f4

      - id: side_padding
        type: f4

      - id: is_vertical
        type: b1

      - id: max_value
        type: f4

      - id: min_value
        type: f4

      - id: init_value
        type: f4

      - id: down
        type: code

      - id: up
        type: code

      - id: scroll
        type: code
