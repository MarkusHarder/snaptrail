import { Component, input } from '@angular/core';
import { Thumbnail } from '../../models/session';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-thumbnail-detail-display',
  imports: [CommonModule],
  templateUrl: './thumbnail-detail-display.component.html',
  styleUrl: './thumbnail-detail-display.component.css',
})
export class ThumbnailDetailDisplayComponent {
  thumbnail = input.required<Thumbnail>();
}
