import { Component, Output, EventEmitter } from '@angular/core';
import { Game } from '../../../models/game.model';
import { GameService } from '../../../services/game.service';

@Component({
  selector: 'app-add-form',
  templateUrl: './add-form.component.html',
  styleUrls: ['./add-form.component.scss']
})
export class AddFormComponent {
  game: Game = {}; 
  displayDialog: boolean = true; 

  constructor(private gameService:GameService){}

  @Output() closeForm = new EventEmitter<void>(); 

  saveGame(): void {
    const data = {
      title: this.game.title,
      genre: this.game.genre,
      developer: this.game.developer,
      platform: this.game.platform,
      price: this.game.price,
      stock: this.game.stock
    };

    this.gameService.create(data)
      .subscribe({
        next: (res) => {
          console.log(res);
        },
        error: (e) => console.error(e)
      });
  }

  addGame(): void {
    console.log("Game added", this.game); 

    this.saveGame()

    this.game = {}; 
    this.displayDialog = false; 
    this.closeForm.emit(); 
  }

  
  cancel(): void {
    this.displayDialog = false; 
    this.closeForm.emit(); 
  }
}
