import { Component, Input } from '@angular/core';
import { CardData } from '../../models/card-data.model';

@Component({
  selector: 'app-card',
  templateUrl: './card.component.html',
  styleUrl: './card.component.scss'
})
export class CardComponent {

  @Input() singleCard!: CardData;

}
