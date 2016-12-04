#version 330 core

uniform sampler2D tex;
uniform mat4 u_ProjTrans;
uniform mat4 u_ObjTrans;

uniform vec3 u_LightPos;

in vec2 v_TexCoord;
in vec3 v_VertNormal;
in vec3 v_Position;
out vec4 color;

void main(){
	color = texture(tex, v_TexCoord) * 
			(abs(dot(normalize(u_LightPos - v_Position),
					 normalize(v_VertNormal)))* 0.7 + 0.3);
}