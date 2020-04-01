package signer_test

import (
	"fmt"
	"github.com/niclabs/hsm-tools/signer"
	"io"
	"testing"
	"time"
)

type vFile struct {
	data    []byte
	pointer uint
}

func (f *vFile) Read(p []byte) (n int, err error) {
	curData := f.data[f.pointer:]
	if len(curData) <= len(p) {
		copy(p[:len(curData)], curData)
		f.pointer = uint(len(f.data))
		return len(curData), io.EOF
	} else {
		copy(p, curData[:len(p)])
		f.pointer += uint(len(p))
		return len(p), nil
	}
}

func (f *vFile) Write(p []byte) (n int, err error) {
	f.data = append(f.data[:f.pointer], p...)
	return len(p), nil
}

func (f *vFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		if offset < 0 {
			return 0, fmt.Errorf("negative seek")
		}
		f.pointer = uint(offset)
	case io.SeekEnd:
		lenData := int64(len(f.data))
		if offset < 0 && offset > int64(-lenData) {
			f.pointer = uint(lenData + offset)
		}
	case io.SeekCurrent:
		if int64(f.pointer)+offset >= 0 && int64(f.pointer)+offset < int64(len(f.data)) {
			f.pointer += uint(offset)
		}
	}
	return int64(f.pointer), nil
}

func TestSession_FileRSASign(t *testing.T) {
	ctx := &signer.Context{
		ContextConfig: &signer.ContextConfig{
			Zone:       zone,
			CreateKeys: true,
			NSEC3:      false,
			OptOut:     false,
		},
		SignAlgorithm: signer.RSA_SHA256,
		Log:           Log,
	}
	zsk := &vFile{data: []byte(RSAZSK)}
	ksk := &vFile{data: []byte(RSAKSK)}
	session, err := ctx.NewFileSession(zsk, ksk)
	if err != nil {
		t.Errorf("%s", err)
		return
	}
	out, err := sign(t, ctx, session)
	if err != nil {
		t.Errorf("signing failed: %s", err)
		return
	}
	defer out.Close()
	if err := signer.VerifyFile(zone, out, Log); err != nil {
		t.Errorf("Error verifying output: %s", err)
	}
	return
}

func TestSession_FileRSASignNSEC3(t *testing.T) {
	ctx := &signer.Context{
		ContextConfig: &signer.ContextConfig{
			Zone:       zone,
			CreateKeys: true,
			NSEC3:      true,
			OptOut:     false,
		},
		SignAlgorithm: signer.RSA_SHA256,
		Log:           Log,
	}
	zsk := &vFile{data: []byte(RSAZSK)}
	ksk := &vFile{data: []byte(RSAKSK)}
	session, err := ctx.NewFileSession(zsk, ksk)
	if err != nil {
		t.Errorf("%s", err)
		return
	}
	out, err := sign(t, ctx, session)
	if err != nil {
		t.Errorf("signing failed: %s", err)
		return
	}
	defer out.Close()
	if err := signer.VerifyFile(zone, out, Log); err != nil {
		t.Errorf("Error verifying output: %s", err)
	}
	return
}

func TestSession_FileRSASignNSEC3OptOut(t *testing.T) {
	ctx := &signer.Context{
		ContextConfig: &signer.ContextConfig{
			Zone:       zone,
			CreateKeys: true,
			NSEC3:      true,
			OptOut:     true,
		},
		SignAlgorithm: signer.RSA_SHA256,
		Log:           Log,
	}
	zsk := &vFile{data: []byte(RSAZSK)}
	ksk := &vFile{data: []byte(RSAKSK)}
	session, err := ctx.NewFileSession(zsk, ksk)
	if err != nil {
		t.Errorf("%s", err)
		return
	}
	out, err := sign(t, ctx, session)
	if err != nil {
		t.Errorf("signing failed: %s", err)
		return
	}
	defer out.Close()
	if err := signer.VerifyFile(zone, out, Log); err != nil {
		t.Errorf("Error verifying output: %s", err)
	}
	return
}

func TestSession_FileECDSASign(t *testing.T) {
	ctx := &signer.Context{
		ContextConfig: &signer.ContextConfig{
			Zone:       zone,
			CreateKeys: true,
			NSEC3:      false,
			OptOut:     false,
		},
		SignAlgorithm: signer.ECDSA_P256_SHA256,
		Log:           Log,
	}
	zsk := &vFile{data: []byte(ECZSK)}
	ksk := &vFile{data: []byte(ECKSK)}
	session, err := ctx.NewFileSession(zsk, ksk)
	if err != nil {
		t.Errorf("%s", err)
		return
	}
	out, err := sign(t, ctx, session)
	if err != nil {
		t.Errorf("signing failed: %s", err)
		return
	}
	defer out.Close()
	if err := signer.VerifyFile(zone, out, Log); err != nil {
		t.Errorf("Error verifying output: %s", err)
	}
	return
}

func TestSession_FileECDSASignNSEC3(t *testing.T) {
	ctx := &signer.Context{
		ContextConfig: &signer.ContextConfig{
			Zone:       zone,
			CreateKeys: true,
			NSEC3:      true,
			OptOut:     false,
		},
		SignAlgorithm: signer.ECDSA_P256_SHA256,
		Log:           Log,
	}
	zsk := &vFile{data: []byte(ECZSK)}
	ksk := &vFile{data: []byte(ECKSK)}
	session, err := ctx.NewFileSession(zsk, ksk)
	if err != nil {
		t.Errorf("%s", err)
		return
	}
	out, err := sign(t, ctx, session)
	if err != nil {
		t.Errorf("signing failed: %s", err)
		return
	}
	defer out.Close()
	if err := signer.VerifyFile(zone, out, Log); err != nil {
		t.Errorf("Error verifying output: %s", err)
	}
	return
}

func TestSession_FileECDSASignNSEC3OptOut(t *testing.T) {
	ctx := &signer.Context{
		ContextConfig: &signer.ContextConfig{
			Zone:       zone,
			CreateKeys: true,
			NSEC3:      true,
			OptOut:     true,
		},
		SignAlgorithm: signer.ECDSA_P256_SHA256,
		Log:           Log,
	}
	zsk := &vFile{data: []byte(ECZSK)}
	ksk := &vFile{data: []byte(ECKSK)}
	session, err := ctx.NewFileSession(zsk, ksk)
	if err != nil {
		t.Errorf("%s", err)
		return
	}
	out, err := sign(t, ctx, session)
	if err != nil {
		t.Errorf("signing failed: %s", err)
		return
	}
	defer out.Close()
	if err := signer.VerifyFile(zone, out, Log); err != nil {
		t.Errorf("Error verifying output: %s", err)
	}
	return
}

func TestSession_FileExpiredSig(t *testing.T) {
	ctx := &signer.Context{
		ContextConfig: &signer.ContextConfig{
			Zone:       zone,
			CreateKeys: true,
			NSEC3:      false,
			OptOut:     false,
		},
		Log:           Log,
		SignAlgorithm: signer.ECDSA_P256_SHA256,
		SignExpDate:   time.Now().AddDate(-1, 0, 0),
	}
	zsk := &vFile{data: []byte(ECZSK)}
	ksk := &vFile{data: []byte(ECKSK)}
	session, err := ctx.NewFileSession(zsk, ksk)
	if err != nil {
		t.Errorf("%s", err)
		return
	}
	out, err := sign(t, ctx, session)
	if err != nil {
		t.Errorf("signing failed: %s", err)
		return
	}
	defer out.Close()
	if err := signer.VerifyFile(zone, out, Log); err == nil {
		t.Errorf("output should be alerted as expired, but it was not")
		return
	}
	return
}