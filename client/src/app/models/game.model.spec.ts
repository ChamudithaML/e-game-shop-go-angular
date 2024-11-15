import { Game } from './game.model';

describe('Game', () => {
  it('should define a Game structure', () => {
    const game: Game = {
      id: 1,
      title: 'Sample Game',
      genre: 'Action',
      developer: 'Dev Studio',
      platform: 'PC',
      price: 59.99,
      stock: 100
    };

    expect(game.id).toBeDefined();
    expect(game.title).toBeDefined();
    expect(game.genre).toBeDefined();
    expect(game.developer).toBeDefined();
    expect(game.platform).toBeDefined();
    expect(game.price).toBeDefined();
    expect(game.stock).toBeDefined();
  });
});
