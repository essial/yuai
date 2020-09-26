package yuai

import (
	"log"

	"github.com/hajimehoshi/ebiten"
)

type HBox struct {
	desktop          *Desktop
	children         []Widget
	dirty            bool
	padding          int
	childSpacing     int
	expandChild      bool
	stretchComponent Widget
	alignment        HAlign
}

func (d *Desktop) CreateHBox() *HBox {
	result := &HBox{
		desktop:      d,
		children:     []Widget{},
		dirty:        false,
		padding:      0,
		childSpacing: 4,
		expandChild:  false,
		alignment:    HAlignLeft,
	}

	return result
}

func (h *HBox) Render(screen *ebiten.Image, x, y, width, height int) {
	if width <= 0 || height <= 0 {
		return
	}

	visibleChildren := 0
	for idx := range h.children {
		cw, ch := h.children[idx].GetRequestedSize()
		if cw <= 0 || ch <= 0 {
			continue
		}
		visibleChildren++
	}

	var childWidth int

	totalChildWidth := 0

	if !h.expandChild {
		for idx := range h.children {
			childWidth, _ = h.children[idx].GetRequestedSize()
			totalChildWidth += childWidth
		}
		totalChildWidth += (visibleChildren - 1) * h.childSpacing
	}

	curY := y + h.padding
	curX := 0
	extraWidth := 0

	if h.expandChild {
		curX = x + h.padding
	} else {
		switch h.alignment {
		case HAlignLeft:
			curX = x + h.padding
		case HAlignCenter:
			curX = x + (width / 2) - (totalChildWidth / 2)
		case HAlignRight:
			curY = y + height - totalChildWidth
		default:
			log.Fatal("unknown HAlign type specified")
		}

		extraWidth = width - totalChildWidth
	}

	if h.expandChild {
		childWidth = (width - (h.padding * 2) - ((visibleChildren - 1) * h.childSpacing)) / visibleChildren
	}

	for idx := range h.children {
		if !h.expandChild {
			childWidth, _ = h.children[idx].GetRequestedSize()
			if h.alignment == HAlignLeft && h.stretchComponent == h.children[idx] {
				childWidth += extraWidth
			}
		} else {
			cw, _ := h.children[idx].GetRequestedSize()
			if cw <= 0 {
				continue
			}
		}

		if childWidth <= 0 {
			continue
		}

		h.children[idx].Render(screen, curX, curY, childWidth, height-(h.padding*2))
		curX += childWidth + h.childSpacing
	}
}

func (h *HBox) Update() (dirty bool) {
	if h.dirty {
		h.Invalidate()
	}

	dirty = false
	for idx := range h.children {
		childDirty := h.children[idx].Update()

		if childDirty {
			dirty = true
		}
	}

	if dirty {
		h.dirty = true
	}

	return dirty
}

func (h *HBox) GetRequestedSize() (int, int) {
	tw := 0
	th := 0

	for idx := range h.children {
		cw, ch := h.children[idx].GetRequestedSize()
		if th < ch {
			th = ch
		}
		tw += cw
	}

	return tw, th
}

func (h *HBox) Invalidate() {
	for idx := range h.children {
		h.children[idx].Invalidate()
	}
}

func (h *HBox) AddChild(widget Widget) {
	h.children = append(h.children, widget)
	h.dirty = true
}

func (h *HBox) SetAlignment(align HAlign) {
	h.alignment = align
	h.dirty = true
}

func (h *HBox) GetAlignment() HAlign {
	return h.alignment
}

func (h *HBox) SetChildSpacing(spacing int) {
	h.childSpacing = spacing
	h.dirty = true
}

func (h *HBox) GetChildSpacing() int {
	return h.childSpacing
}

func (h *HBox) SetPadding(padding int) {
	h.padding = padding
	h.dirty = true
}

func (h *HBox) GetPadding() int {
	return h.padding
}

func (h *HBox) SetExpandChild(expand bool) {
	h.expandChild = expand
	h.dirty = true
}

func (h *HBox) GetExpandChild() bool {
	return h.expandChild
}

func (h *HBox) SetStretchComponent(widget Widget) {
	h.stretchComponent = widget
}
