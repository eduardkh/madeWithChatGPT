import pygame
import sys

# --- Config ---
WIDTH = 800
HEIGHT = 600
FPS = 60

PADDLE_WIDTH = 120
PADDLE_HEIGHT = 15
BALL_RADIUS = 10
PADDLE_SPEED = 7
BALL_SPEED = 5

# --- Init ---
pygame.init()
screen = pygame.display.set_mode((WIDTH, HEIGHT))
pygame.display.set_caption("Breakout - Scaffold")
clock = pygame.time.Clock()

# --- Objects ---
paddle = pygame.Rect(
    WIDTH // 2 - PADDLE_WIDTH // 2,
    HEIGHT - 50,
    PADDLE_WIDTH,
    PADDLE_HEIGHT
)

ball = pygame.Rect(
    WIDTH // 2,
    HEIGHT // 2,
    BALL_RADIUS * 2,
    BALL_RADIUS * 2
)

ball_dx = BALL_SPEED
ball_dy = -BALL_SPEED

# --- Game Loop ---
while True:
    clock.tick(FPS)

    # --- Events ---
    for event in pygame.event.get():
        if event.type == pygame.QUIT:
            pygame.quit()
            sys.exit()

    # --- Input ---
    keys = pygame.key.get_pressed()
    if keys[pygame.K_LEFT] and paddle.left > 0:
        paddle.x -= PADDLE_SPEED
    if keys[pygame.K_RIGHT] and paddle.right < WIDTH:
        paddle.x += PADDLE_SPEED

    # --- Ball Movement ---
    ball.x += ball_dx
    ball.y += ball_dy

    # --- Wall Collision ---
    if ball.left <= 0 or ball.right >= WIDTH:
        ball_dx *= -1
    if ball.top <= 0:
        ball_dy *= -1

    # --- Paddle Collision ---
    if ball.colliderect(paddle):
        ball_dy *= -1

    # --- Ball falls below ---
    if ball.bottom >= HEIGHT:
        pygame.quit()
        sys.exit()

    # --- Draw ---
    screen.fill((20, 20, 20))
    pygame.draw.rect(screen, (200, 200, 200), paddle)
    pygame.draw.ellipse(screen, (255, 100, 100), ball)
    pygame.display.flip()
