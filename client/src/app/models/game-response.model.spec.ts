import { GameResponse } from './game-response.model';

describe('GameResponse', () => {
  it('should define a GameResponse structure', () => {
    const response: GameResponse = {
      status: 200,
      message: 'Success',
      data: {}
    };

    expect(response.status).toBeDefined();
    expect(response.message).toBeDefined();
    expect(response.data).toBeDefined();
  });
});
