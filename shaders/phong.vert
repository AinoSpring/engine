#version 410 core

layout (location = 0) in vec3 position;
layout (location = 1) in vec3 normal;
layout (location = 2) in vec2 texCoord;

uniform mat4 model;
uniform mat4 view;
uniform mat4 project;

uniform vec3 lightPos;

out vec3 Normal;
out vec3 FragPos;
out vec3 LightPos;
out vec2 TexCoord;

void main()
{
  TexCoord = texCoord;

  gl_Position = project * view * model * vec4(position, 1.0);

  FragPos = vec3(view * model * vec4(position, 1.0));
  LightPos = vec3(view * vec4(lightPos, 1.0));

  mat3 normMatrix = mat3(transpose(inverse(view))) * mat3(transpose(inverse(model)));
  Normal = normMatrix * normal;
}
