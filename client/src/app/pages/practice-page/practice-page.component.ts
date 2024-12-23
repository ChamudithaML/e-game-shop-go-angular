import { Component } from '@angular/core';
import { DialogService, DynamicDialogRef } from 'primeng/dynamicdialog';
import { TempDialogComponent } from '../../components/practice-components/temp-dialog/temp-dialog.component';

@Component({
  selector: 'app-practice-page',
  templateUrl: './practice-page.component.html',
  styleUrl: './practice-page.component.scss',
})
export class PracticePageComponent {

  ref: DynamicDialogRef | undefined

  constructor(private dialogService: DialogService) {}

  openDialog() {
    this.ref = this.dialogService.open(TempDialogComponent, {
     
      width: '800px',
      
      contentStyle: { "max-height": "800px", "overflow": "auto" }
    });

    this.ref.onClose.subscribe((result: any) => {
      if (result) {
        console.log('Dialog closed with result:', result);
      }
    });
  }


}
