package main

import rl "github.com/gen2brain/raylib-go/raylib"

type SunRay struct {
	mouse     *Mouse
	effect    *Effect
	start     []*rl.Vector2
	end       *rl.Vector2
	color     *rl.Color
	total     int32
	divisor   int32
	thickness float32
	opacity   float32
}

// ---------------------------------------------------------------------------------------------------------------------

func (s *SunRay) init() {
	for i := range s.effect.totalParticles {
		if i%s.divisor == 0 {
			s.total++
		}
	}

	s.color = &rl.Color{R: 255, G: 255, B: 255, A: uint8(255 * s.opacity)}
}

// ---------------------------------------------------------------------------------------------------------------------

func (s *SunRay) calculateLocations() {
	s.start = make([]*rl.Vector2, 0)
	for i := range s.effect.totalParticles {
		if i%s.divisor == 0 {
			s.start = append(
				s.start,
				&rl.Vector2{
					X: float32(s.effect.particles[i].x),
					Y: float32(s.effect.particles[i].y),
				},
			)
		}
	}

	s.end = &rl.Vector2{
		X: float32(s.mouse.x),
		Y: float32(s.mouse.y),
	}
}

// ---------------------------------------------------------------------------------------------------------------------

func (s *SunRay) draw() {
	for i := range s.total {
		rl.DrawLineEx(
			*s.start[i],
			*s.end,
			s.thickness,
			*s.color,
		)
	}
}

// ---------------------------------------------------------------------------------------------------------------------
