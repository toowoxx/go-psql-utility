FROM archlinux

RUN pacman -Syu --noconfirm --needed go postgresql

ENTRYPOINT ["/src/test.sh"]
