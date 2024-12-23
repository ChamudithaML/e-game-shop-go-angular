
import { Component, Output, EventEmitter } from '@angular/core';
// import { Game } from '../../models/game.model';
// import { GameService } from '../../services/game.service';

import { FormGroup, FormControl, Validators } from '@angular/forms';

@Component({
  selector: 'app-temp-dialog',
  templateUrl: './temp-dialog.component.html',
  styleUrl: './temp-dialog.component.scss'
})

export class TempDialogComponent {

  form: FormGroup;

  constructor() {
    this.form = new FormGroup({
      firstname: new FormControl('', [Validators.required, Validators.minLength(3)]),
      lastname: new FormControl('', [Validators.required, Validators.minLength(3)]),
      address: new FormControl('', [Validators.required, Validators.minLength(10)]),
      city: new FormControl('', [Validators.required]),
      state: new FormControl('', [Validators.required]),
      zip: new FormControl('', [Validators.required, Validators.pattern('^[0-9]{5}$')]) // 5-digit zip code validation
    });
  }

  // Helper method to check if a form control is invalid and has been touched
  isInvalid(controlName: string): boolean {
    const control = this.form.get(controlName);
    return !!control && control.invalid && control.touched;
  }

  onSubmit(): void {
    if (this.form.valid) {
      console.log("Form Submitted:", this.form.value);
    } else {
      // Mark all fields as touched to show validation messages
      this.form.markAllAsTouched();
    }
  }

}


































// game: Game = {}; 
// displayDialog: boolean = true; 

// constructor(private gameService:GameService){}

// @Output() closeForm = new EventEmitter<void>(); 

// saveGame(): void {
//   const data = {
//     title: this.game.title,
//     genre: this.game.genre,
//     developer: this.game.developer,
//     platform: this.game.platform,
//     price: this.game.price,
//     stock: this.game.stock
//   };

//   this.gameService.create(data)
//     .subscribe({
//       next: (res) => {
//         console.log(res);
//       },
//       error: (e) => console.error(e)
//     });
// }

// addGame(): void {
//   console.log("Game added", this.game); 

//   this.saveGame()

//   this.game = {}; 
//   this.displayDialog = false; 
//   this.closeForm.emit(); 
// }

// cancel(): void {
//   this.displayDialog = false; 
//   this.closeForm.emit(); 
// }