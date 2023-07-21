#version 410 core

in vec2 TexCoord;

out vec4 color;

uniform sampler2D bg_texture;
uniform vec3 objectColor;
uniform vec3 lightColor;

void main()
{
    // the color of the light "reflects" off the object
    color = vec4(texture(bg_texture, TexCoord) * lightColor, 1.0f);
}
