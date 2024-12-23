import { Component, OnInit } from '@angular/core';
import { Game } from '../../models/game.model';
import { GameService } from '../../services/game.service';
import { DialogService, DynamicDialogRef } from 'primeng/dynamicdialog';
import { AddEditFormComponent } from '../Forms/add-edit-form/add-edit-form.component';

@Component({
  selector: 'app-games',
  templateUrl: './games.component.html',
  styleUrls: ['./games.component.scss']
})
export class GamesComponent implements OnInit {

  games: Game[] = [];
  title = ''
  displayAddForm: boolean = false;
  displayEditForm: boolean = false;
  selectedGame?: Game;

  ref: DynamicDialogRef | undefined

  constructor(private gameService: GameService, private dialogService: DialogService) { }

  ngOnInit(): void {
    this.retrieveGames();
  }

  // retrieveGames(): void {
  //   this.gameService.getAll()
  //     .subscribe({
  //       next: (data) => {
  //         this.games = data.data;
  //         console.log(this.games);
  //       },
  //       error: (e) => console.error(e)
  //     });
  // }


  retrieveGames(): void {
    this.gameService.getAll().subscribe(response => {
      this.games = response.data.data;
    }, error => {
      console.error("Error retrieving games:", error);
    });
  }

  addGame(): void {
    console.log("Open add popup")
    this.displayAddForm = true
  }

  reloadData(): void {
    this.displayAddForm = false
    this.retrieveGames()
  }

  editGame(game: Game): void {
    console.log("Open edit popup")
    this.selectedGame = { ...game }
    this.displayEditForm = true
  }

  deleteGame(id: any): void {
    this.gameService.delete(id)
      .subscribe({
        next: (res) => {
          console.log(res);
        },
        error: (e) => console.error(e)
      });

    console.log("sde")
    this.retrieveGames()
  }

  openDialog() {
    this.ref = this.dialogService.open(AddEditFormComponent, {

      width: '800px',

      contentStyle: { "max-height": "800px", "overflow": "auto" }
    });

    this.ref.onClose.subscribe((result: any) => {
      if (result && result.updated) {
        console.log('Dialog closed with result:', result);
        this.retrieveGames();
      }
    });
  }

}
