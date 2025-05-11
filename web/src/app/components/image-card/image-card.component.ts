import { Component, input } from '@angular/core';
import { Thumbnail } from '../../models/session';
import { ImageModule } from 'primeng/image';
import { ThumbnailDetailDisplayComponent } from '../thumbnail-detail-display/thumbnail-detail-display.component';

@Component({
  selector: 'app-image-card',
  imports: [ThumbnailDetailDisplayComponent, ImageModule],
  templateUrl: './image-card.component.html',
  styleUrl: './image-card.component.css',
})
export class ImageCardComponent {
  thumbnail = input.required<Thumbnail>();
  displayDetails = input<boolean>(false);
}
