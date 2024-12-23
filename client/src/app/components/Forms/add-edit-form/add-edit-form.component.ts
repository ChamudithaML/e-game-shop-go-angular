import { Component } from '@angular/core';
import { FormGroup, FormControl, Validators } from '@angular/forms';
import { GameService } from '../../../services/game.service';
import { Game } from '../../../models/game.model';
import { DynamicDialogRef } from 'primeng/dynamicdialog';

@Component({
  selector: 'app-add-edit-form',
  templateUrl: './add-edit-form.component.html',
  styleUrl: './add-edit-form.component.scss'
})
export class AddEditFormComponent {

  gameForm: FormGroup;
  game?: Game;

  constructor(private gameService: GameService, private dialogRef: DynamicDialogRef) {

    this.gameForm = new FormGroup({
      title: new FormControl('', [Validators.required, Validators.minLength(3)]),
      genre: new FormControl('', [Validators.required, Validators.minLength(3)]),
      developer: new FormControl('', [Validators.required, Validators.minLength(3)]),
      platform: new FormControl('', [Validators.required, Validators.minLength(3)]),
      price: new FormControl('', [Validators.required, Validators.min(0)]),
      stock: new FormControl('', [Validators.required, Validators.min(0)])
    });
  }

  isInvalid(controlName: string): boolean {
    const control = this.gameForm.get(controlName);
    return !!control && control.invalid && (control.dirty || control.touched);
  }

  onSubmit(): void {
    if (this.gameForm.valid) {
      this.game = {
        title: this.gameForm.value.title,
        genre: this.gameForm.value.genre,
        developer: this.gameForm.value.developer,
        platform: this.gameForm.value.platform,
        price: this.gameForm.value.price,
        stock: this.gameForm.value.stock
      };

    }

    if (this.game) {
      this.gameService.create(this.game)
        .subscribe({
          next: (res) => {
            console.log(res);
          },
          error: (e) => console.error(e)
        });
    } else {
      console.error('Game is undefined!');
    }

    this.dialogRef.close({ updated: true });


  }

  onCancel(): void {
    this.dialogRef.close();
  }

  

}


// this.gameService.create(data)
//       .subscribe({
//         next: (res) => {
//           console.log(res);
//         },
//         error: (e) => console.error(e)
//       });