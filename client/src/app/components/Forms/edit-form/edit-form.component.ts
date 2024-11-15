import { Component, Output, EventEmitter, Input, SimpleChanges, OnChanges  } from '@angular/core';
import { Game } from '../../../models/game.model';
import { GameService } from '../../../services/game.service';

@Component({
  selector: 'app-edit-form',
  templateUrl: './edit-form.component.html',
  styleUrl: './edit-form.component.scss'
})
export class EditFormComponent  {
  @Output() closeForm = new EventEmitter<void>(); 
  @Input() gameIn?: Game;

  game: Game = {}; 
  displayDialog: boolean = true; 

  constructor(private gameService:GameService){}

  ngOnChanges(changes: SimpleChanges) {
    if ('gameIn' in changes) {
      this.game = { ...this.gameIn };
    }
  }

  updateGame(): void {
    const data = {
      id:this.game.id,
      title: this.game.title,
      genre: this.game.genre,
      developer: this.game.developer,
      platform: this.game.platform,
      price: this.game.price,
      stock: this.game.stock
    };

    this.gameService.update(data.id, data)
      .subscribe({
        next: (res) => {
          console.log(res);
        },
        error: (e) => console.error(e)
      });
  }

  editGame(): void {
    console.log("Game edited", this.game); 

    this.updateGame()

    this.game = {}; 
    this.displayDialog = false; 
    this.closeForm.emit(); 
  }

  
  cancel(): void {
    this.displayDialog = false; 
    this.closeForm.emit(); 
  }
}

