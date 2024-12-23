import { Component } from '@angular/core';
import { CardData } from '../../models/card-data.model';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrl: './home.component.scss'
})
export class HomeComponent {

  cardsData: CardData[] = [
    {
      title: "The Legend of Zelda: Breath of the Wild",
      description: "An open-world action-adventure game set in a vast kingdom with endless exploration and epic quests.",
      img_url: "assets/images/img.jpeg"
    },
    {
      title: "Elden Ring",
      description: "A fantasy action RPG combining a rich story and challenging combat, set in a dark and mysterious world.",
      img_url: "assets/images/img.jpeg"
    },
    {
      title: "Minecraft",
      description: "A sandbox game where you can build, explore, and survive in a blocky, pixelated world.",
      img_url: "assets/images/img.jpeg"
    },
    {
      title: "Red Dead Redemption 2",
      description: "An action-adventure game set in the American Wild West, featuring an open-world experience with a rich storyline.",
      img_url: "assets/images/img.jpeg"
    },
    {
      title: "The Witcher 3: Wild Hunt",
      description: "An open-world RPG where you play as Geralt of Rivia, a monster hunter navigating through a war-torn world.",
      img_url: "assets/images/img.jpeg"
    },
    {
      title: "Horizon Zero Dawn",
      description: "An action RPG set in a post-apocalyptic world, where robotic creatures roam and humanity is on the brink of extinction.",
      img_url: "assets/images/img.jpeg"
    },
    {
      title: "Assassin's Creed Valhalla",
      description: "An action RPG set in the Viking age, where you play as a Viking raider exploring and conquering England.",
      img_url: "assets/images/img.jpeg"
    }
  ];

  myFunc(): void {
  }

}
