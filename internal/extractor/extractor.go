package extractor

import (
	"fmt"
	"snaptrail/internal/structs"

	"github.com/dsoprea/go-exif"
	"github.com/rs/zerolog/log"
)

type ValueType int

type TagMapping struct {
	exifTagName string
	path        string
	targetField string
}

func CreateExifMetadataFromImage(thumbnail []byte) (md *structs.ExifMetadata, err error) {
	md = &structs.ExifMetadata{}
	rawExif, err := exif.SearchAndExtractExif(thumbnail)
	if err != nil {
		log.Err(err).Msg("failed to search exif data for thumbnail")
		return nil, err
	}
	im := exif.NewIfdMapping()

	err = exif.LoadStandardIfds(im)
	if err != nil {
		log.Err(err).Msg("failed to load standard idfs")
		return nil, err
	}

	ti := exif.NewTagIndex()

	_, index, err := exif.Collect(im, ti, rawExif)
	if err != nil {
		log.Err(err).Msg("failed to collect index")
		return nil, err
	}

	tagMapping := []TagMapping{
		{"Model", exif.IfdPathStandard, "CameraModel"},
		{"Make", exif.IfdPathStandard, "Make"},
		{"DateTime", exif.IfdPathStandard, "DateTime"},
		{"FNumber", exif.IfdPathStandardExif, "Aperture"},
		{"ISOSpeedRatings", exif.IfdPathStandardExif, "ISO"},
		{"LensModel", exif.IfdPathStandardExif, "Lens"},
		{"FocalLength", exif.IfdPathStandardExif, "FocalLength"},
		{"ExposureTime", exif.IfdPathStandardExif, "Exposure"},
	}

	for _, mapping := range tagMapping {
		ifd, err := getIfd(index.RootIfd, mapping.path)
		if err != nil {
			return nil, err
		}
		result, err := findTag(ifd, mapping.exifTagName)
		if err != nil {
			return nil, err
		}

		rawValue, err := ifd.TagValue(result)
		if err != nil {
			return nil, err
		}
		err = parseAndAssignTagValue(md, rawValue, mapping.targetField)
		if err != nil {
			return nil, err
		}
	}

	return md, nil
}

func getIfd(rootIfd *exif.Ifd, path string) (*exif.Ifd, error) {
	if path == "IFD" {
		return rootIfd, nil
	}
	return rootIfd.ChildWithIfdPath(path)
}

func findTag(ifd *exif.Ifd, tagName string) (*exif.IfdTagEntry, error) {
	results, err := ifd.FindTagWithName(tagName)
	if err != nil || len(results) == 0 {
		return nil, fmt.Errorf("tag %s not found", tagName)
	}
	return results[0], nil
}

func parseAndAssignTagValue(md *structs.ExifMetadata, rawValue any, target string) error {
	log.Info().Msgf("going for target: %s\n", target)
	switch target {
	case "CameraModel":
		s, ok := rawValue.(string)
		if !ok {
			return fmt.Errorf("unable to parse CameraModel as string")
		}
		md.CameraModel = s
		log.Info().Msgf("assigned %s to %s\n", s, target)
		return nil
	case "Make":
		s, ok := rawValue.(string)
		if !ok {
			return fmt.Errorf("unable to parse Make as string")
		}
		md.Make = s
		log.Info().Msgf("assigned %s to %s\n", s, target)
		return nil
	case "DateTime":
		s, ok := rawValue.(string)
		if !ok {
			return fmt.Errorf("unable to parse DateTime as string")
		}
		md.DateTime = s
		log.Info().Msgf("assigned %s to %s\n", s, target)
		return nil
	case "Aperture":
		switch v := rawValue.(type) {
		case exif.Rational:
			md.Aperture = rationalToFloat(v)
			log.Info().Msgf("assigned %f to %s\n", md.Aperture, target)
			return nil
		case []exif.Rational:
			md.Aperture = rationalSliceToFloat(v)
			return nil
		default:
			return fmt.Errorf("unable to parse Aperture as float")
		}
	case "ISO":
		num, err := parseInt(rawValue)
		if err != nil {
			return fmt.Errorf("unable to parse ISOSPeedRatings as int")
		}
		md.ISO = num
		log.Info().Msgf("assigned %d to %s\n", num, target)
		return nil
	case "FocalLength":
		switch v := rawValue.(type) {
		case exif.Rational:
			md.FocalLength = rationalToFloat(v)
			log.Info().Msgf("assigned %f to %s\n", md.FocalLength, target)
			return nil
		case []exif.Rational:
			md.FocalLength = rationalSliceToFloat(v)
			return nil
		default:
			return fmt.Errorf("unable to parse FocalLength as float")
		}
	case "Exposure":
		switch v := rawValue.(type) {
		case exif.Rational:
			md.Exposure = rationalToString(v)
			log.Info().Msgf("assigned %s to %s\n", md.Exposure, target)
			return nil
		case []exif.Rational:
			md.Exposure = rationalSliceToString(v)
			return nil
		default:
			return fmt.Errorf("unable to parse Exposure as string")
		}
	case "Lens":
		s, ok := rawValue.(string)
		if !ok {
			return fmt.Errorf("unable to parse Lens as string")
		}
		md.LensModel = s
		log.Info().Msgf("assigned %s to %s\n", s, target)
	}
	return nil
}

func rationalSliceToFloat(ratios []exif.Rational) float64 {
	if len(ratios) > 0 {
		return rationalToFloat(ratios[0])
	}
	return 0.0
}

func rationalSliceToString(ratios []exif.Rational) string {
	if len(ratios) > 0 {
		return rationalToString(ratios[0])
	}
	return ""
}

func rationalToFloat(r exif.Rational) float64 {
	return float64(r.Numerator) / float64(r.Denominator)
}

func rationalToString(r exif.Rational) string {
	return fmt.Sprintf("%d/%d", r.Numerator, r.Denominator)
}

func parseInt(raw any) (int, error) {
	switch v := raw.(type) {
	case int:
		return v, nil
	case int64:
		return int(v), nil
	case uint16:
		return int(v), nil
	case uint32:
		return int(v), nil
	case []uint16:
		if len(v) > 0 {
			return int(v[0]), nil
		}
	case []uint32:
		if len(v) > 0 {
			return int(v[0]), nil
		}
	case []int:
		if len(v) > 0 {
			return v[0], nil
		}
	case []int64:
		if len(v) > 0 {
			return int(v[0]), nil
		}
	}

	return 0, fmt.Errorf("unsupported int type: %T", raw)
}
