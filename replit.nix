{ pkgs }: {
    deps = [
        pkgs.nodejs-16_x
        pkgs.sudo
        pkgs.go_1_17
        pkgs.gopls
    ];
}